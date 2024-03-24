package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rarimo/rarimo-core/x/cbank/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CommitmentAll(c context.Context, req *types.QueryGetAllCommitmentRequest) (*types.QueryGetAllCommitmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var commitments []types.Commitment

	store := ctx.KVStore(k.storeKey)
	confirmationStore := prefix.NewStore(store, types.KeyPrefix(types.CommitmentKeyPrefix))

	pageRes, err := query.Paginate(confirmationStore, req.Pagination, func(key []byte, value []byte) error {
		var commitment types.Commitment
		if err := k.cdc.Unmarshal(value, &commitment); err != nil {
			return err
		}

		commitments = append(commitments, commitment)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetAllCommitmentResponse{Commitment: commitments, Pagination: pageRes}, nil
}

func (k Keeper) Commitment(c context.Context, req *types.QueryGetCommitmentRequest) (*types.QueryGetCommitmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCommitment(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCommitmentResponse{Commitment: val}, nil
}
