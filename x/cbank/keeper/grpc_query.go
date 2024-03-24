package keeper

import (
	"context"
	"github.com/rarimo/rarimo-core/x/cbank/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CommitmentAll(c context.Context, req *types.QueryGetAllCommitmentRequest) (*types.QueryGetAllCommitmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) Commitment(c context.Context, req *types.QueryGetCommitmentRequest) (*types.QueryGetCommitmentResponse, error) {
	//TODO implement me
	panic("implement me")
}
