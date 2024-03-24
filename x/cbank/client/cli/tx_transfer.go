package cli

import (
	"context"
	"errors"
	"github.com/cloudflare/bn256"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/distributed-lab/bulletproofs"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/rarimo-core/x/cbank/pkg"
	"github.com/rarimo/rarimo-core/x/cbank/types"
	"github.com/spf13/cobra"
	"math/big"
)

func CmdTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [a1] [a2] [v] [B1] [B2] [denom]",
		Short: "Make 1-to-1 transfer for sender size",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			arg_a1 := args[0]
			arg_a2 := args[1]
			arg_v := args[2]
			argB1 := args[3]
			argB2 := args[4]
			argDenom := args[5]

			a1 := new(big.Int).SetBytes(hexutil.MustDecode(arg_a1))
			a2 := new(big.Int).SetBytes(hexutil.MustDecode(arg_a2))

			params, err := queryClient.Params(context.TODO(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			v, ok := new(big.Int).SetString(arg_v, 10)
			if !ok {
				return errors.New("invalid v: should be in dec")
			}

			A1 := pkg.PublicKey(a1, params.Params.G.MustToBN256G1())
			A2 := pkg.PublicKey(a2, params.Params.G.MustToBN256G1())

			comIn := new(bn256.G1).Add(A1, pkg.PublicKey(v, params.Params.HVec[0].MustToBN256G1()))

			B1, err := new(types.Point).Decompress(argB1)
			if err != nil {
				return err
			}

			B2, err := new(types.Point).Decompress(argB2)
			if err != nil {
				return err
			}

			k1 := pkg.HashDH(B1.MustToBN256G1(), a2)
			k2 := pkg.HashDH(B2.MustToBN256G1(), a2)

			comOut := new(bn256.G1).Add(pkg.PublicKey(k1, params.Params.G.MustToBN256G1()), pkg.PublicKey(v, params.Params.HVec[0].MustToBN256G1()))
			addrOut := new(bn256.G1).Add(pkg.PublicKey(k2, params.Params.G.MustToBN256G1()), B2.MustToBN256G1())

			sigCom, err := pkg.SignSchnorr(
				new(big.Int).Mod(new(big.Int).Sub(a1, k1), bn256.Order),
				pkg.PublicKey(new(big.Int).Mod(new(big.Int).Sub(a1, k1), bn256.Order), params.Params.G.MustToBN256G1()),
				params.Params.G.MustToBN256G1(),
				pkg.Msg(comOut.Marshal(), []byte("transfer"), comIn.Marshal()),
			)
			if err != nil {
				return err
			}

			proof := bulletproofs.ProveRange(params.Params.ReciprocalPublic(), bulletproofs.NewKeccakFS(), &bulletproofs.ReciprocalPrivate{
				X:      v,
				M:      bulletproofs.HexMapping(bulletproofs.UInt64Hex(v.Uint64())),
				Digits: bulletproofs.UInt64Hex(v.Uint64()),
				S:      a1,
			})

			sigAddr, err := pkg.SignSchnorr(
				a2,
				A2,
				params.Params.G.MustToBN256G1(),
				pkg.Msg([]byte("transfer address"), comIn.Marshal(), A2.Marshal()),
			)
			if err != nil {
				return err
			}

			msg := types.MsgTransfer{
				In:            []types.Commitment{types.NewCommitment(comIn, A2, argDenom)},
				Out:           []types.Commitment{types.NewCommitment(comOut, addrOut, argDenom)},
				SigCommitment: *new(types.Signature).FromSignature(sigCom),
				SigAddress:    []types.Signature{*new(types.Signature).FromSignature(sigAddr)},
				Proof:         []types.RangeProof{*new(types.RangeProof).FromProof(proof)},
				Creator:       clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
