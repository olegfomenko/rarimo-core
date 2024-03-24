package keeper

import (
	"context"
	"github.com/cloudflare/bn256"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/distributed-lab/bulletproofs"
	"github.com/rarimo/rarimo-core/x/cbank/pkg"
	"github.com/rarimo/rarimo-core/x/cbank/types"
	"math/big"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) Deposit(c context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := m.GetParams(ctx)

	commitmentPub := msg.CommitmentPublicKey.MustToBN256G1()
	address := msg.Address.MustToBN256G1()

	sigMsgHash := pkg.Msg([]byte("deposit"), commitmentPub.Marshal(), address.Marshal())

	if !pkg.VerifySchnorr(msg.SigCommitment.Signature(), msg.CommitmentPublicKey.MustToBN256G1(), params.G.MustToBN256G1(), sigMsgHash) {
		return nil, types.ErrInvalidSignature
	}

	amount, _ := sdk.NewIntFromString(msg.Amount)
	payerAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	moduleAccount := m.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	if err := m.bankKeeper.SendCoins(ctx, payerAddr, moduleAccount.GetAddress(), sdk.NewCoins(sdk.NewCoin(msg.Demon, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error transferring coins: %s", err.Error())
	}

	com := new(bn256.G1).Add(
		new(bn256.G1).ScalarMult(params.HVec[0].MustToBN256G1(), amount.BigInt()),
		commitmentPub,
	)

	commitment := &types.Commitment{
		Commitment: *new(types.Point).MustFromBN256G1(com),
		Address:    msg.Address,
		Denom:      msg.Demon,
	}

	if _, ok := m.GetCommitment(ctx, commitment.Index()); ok {
		return nil, types.ErrCommitmentExist
	}

	m.SetCommitment(ctx, commitment)

	return &types.MsgDepositResponse{}, nil
}

func (m msgServer) Withdraw(c context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := m.GetParams(ctx)

	commitment, ok := m.GetCommitment(ctx, msg.CommitmentIndex)
	if !ok {
		return nil, types.ErrCommitmentNotFound
	}

	amount, _ := sdk.NewIntFromString(msg.Amount)

	aH := new(bn256.G1).ScalarMult(params.HVec[0].MustToBN256G1(), amount.BigInt())
	commitmentPub := new(bn256.G1).Add(
		commitment.Commitment.MustToBN256G1(),
		new(bn256.G1).ScalarMult(aH, big.NewInt(-1)),
	)

	sigMsgHash := pkg.Msg([]byte("withdraw"), commitment.Commitment.MustToBN256G1().Marshal(), commitment.Address.MustToBN256G1().Marshal())

	if !pkg.VerifySchnorr(msg.SigCommitment.Signature(), commitmentPub, params.G.MustToBN256G1(), sigMsgHash) {
		return nil, types.ErrInvalidSignature
	}

	sigMsgHash = pkg.Msg([]byte("withdraw address"), commitment.Commitment.MustToBN256G1().Marshal(), commitment.Address.MustToBN256G1().Marshal())

	if !pkg.VerifySchnorr(msg.SigAddress.Signature(), commitment.Address.MustToBN256G1(), params.G.MustToBN256G1(), sigMsgHash) {
		return nil, types.ErrInvalidSignature
	}

	receiverAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	moduleAccount := m.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	if err := m.bankKeeper.SendCoins(ctx, moduleAccount.GetAddress(), receiverAddr, sdk.NewCoins(sdk.NewCoin(commitment.Denom, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error transferring coins: %s", err.Error())
	}

	m.RemoveCommitment(ctx, msg.CommitmentIndex)

	return &types.MsgWithdrawResponse{}, nil
}

func (m msgServer) Transfer(c context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := m.GetParams(ctx)
	reciprocal := params.ReciprocalPublic()

	var sigPublicKey *bn256.G1

	var sigMsg []byte = make([]byte, 0, (len(msg.In)+len(msg.Out))*64)

	for i, out := range msg.Out {
		if _, ok := m.GetCommitment(ctx, out.Index()); ok {
			return nil, types.ErrCommitmentExist
		}

		if err := bulletproofs.VerifyRange(reciprocal, out.Commitment.MustToBN256G1(), bulletproofs.NewKeccakFS(), msg.Proof[i].Proof()); err != nil {
			return nil, types.ErrInvalidRangeProof
		}

		toAdd := new(bn256.G1).ScalarMult(out.Commitment.MustToBN256G1(), big.NewInt(-1))

		if sigPublicKey == nil {
			sigPublicKey = toAdd
			continue
		}

		sigPublicKey = new(bn256.G1).Add(sigPublicKey, toAdd)

		sigMsg = append(sigMsg, out.Commitment.MustToBN256G1().Marshal()...)
	}

	sigMsg = append(sigMsg, []byte("transfer")...)

	for i, in := range msg.In {
		if _, ok := m.GetCommitment(ctx, in.Index()); !ok {
			return nil, types.ErrCommitmentNotFound
		}

		sigPublicKey = new(bn256.G1).Add(sigPublicKey, in.Commitment.MustToBN256G1())

		sigMsg = append(sigMsg, in.Commitment.MustToBN256G1().Marshal()...)

		sigMsgHash := pkg.Msg([]byte("transfer address"), in.Commitment.MustToBN256G1().Marshal(), in.Address.MustToBN256G1().Marshal())

		if !pkg.VerifySchnorr(msg.SigAddress[i].Signature(), in.Address.MustToBN256G1(), params.G.MustToBN256G1(), sigMsgHash) {
			return nil, types.ErrInvalidSignature
		}
	}

	if !pkg.VerifySchnorr(msg.SigCommitment.Signature(), sigPublicKey, params.G.MustToBN256G1(), pkg.Msg(sigMsg)) {
		return nil, types.ErrInvalidSignature
	}

	for _, in := range msg.In {
		m.RemoveCommitment(ctx, in.Index())
	}

	for _, out := range msg.Out {
		m.SetCommitment(ctx, &out)
	}

	return &types.MsgTransferResponse{}, nil
}
