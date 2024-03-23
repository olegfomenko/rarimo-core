package keeper

import (
	"context"
	"github.com/rarimo/rarimo-core/x/cbank/types"
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
	//TODO implement me
	panic("implement me")
}

func (m msgServer) Withdraw(c context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) Transfer(c context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) FinishTransfer(c context.Context, msg *types.MsgFinishTransfer) (*types.MsgFinishTransferResponse, error) {
	//TODO implement me
	panic("implement me")
}
