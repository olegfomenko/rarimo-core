package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	if err := k.checkCreatorIsValidator(ctx, msg.Creator); err != nil {
		return nil, err
	}

	operation, ok := k.GetOperation(ctx, msg.Operation)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "operation not found")
	}

	k.SetVote(ctx, types.Vote{
		Index: &types.VoteIndex{
			Operation: msg.Operation,
			Validator: msg.Creator,
		},
		Vote: msg.Vote,
	})

	if operation.Status != types.OpStatus_INITIALIZED && operation.Status != types.OpStatus_NOT_APPROVED {
		return &types.MsgVoteResponse{}, nil
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeVoted,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
		sdk.NewAttribute(types.AttributeKeyVotingChoice, msg.Vote.String()),
	))

	yesResult := sdk.ZeroDec()
	noResult := sdk.ZeroDec()

	// Setting votes in validators map
	k.IterateVotes(ctx, msg.Operation, func(vote types.Vote) (stop bool) {
		voter := sdk.MustAccAddressFromBech32(vote.Index.Validator)

		if validator := k.staking.Validator(ctx, sdk.ValAddress(voter)); validator != nil {
			switch vote.Vote {
			case types.VoteType_YES:
				yesResult = yesResult.Add(validator.GetBondedTokens().ToDec())
			case types.VoteType_NO:
				noResult = noResult.Add(validator.GetBondedTokens().ToDec())
			}
		}

		return false
	})

	params := k.GetParams(ctx)
	quorum, _ := sdk.NewDecFromStr(params.VoteQuorum)
	threshold, _ := sdk.NewDecFromStr(params.VoteThreshold)

	totalVotingPower := yesResult.Add(noResult)

	// If there is not enough quorum of votes, finish the flow
	percentVoting := totalVotingPower.Quo(k.staking.TotalBondedTokens(ctx).ToDec())
	if percentVoting.LT(quorum) {
		return &types.MsgVoteResponse{}, nil
	}

	if yesResult.Quo(totalVotingPower).GT(threshold) {
		if err := k.ApproveOperation(ctx, operation); err != nil {
			return nil, err
		}

		return &types.MsgVoteResponse{}, nil
	}

	if err := k.UnapproveOperation(ctx, operation); err != nil {
		return nil, err
	}

	return &types.MsgVoteResponse{}, nil
}

func (k msgServer) UnapproveOperation(ctx sdk.Context, op types.Operation) error {
	op.Status = types.OpStatus_NOT_APPROVED
	k.SetOperation(ctx, op)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationRejected,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
	))

	return nil
}

func (k msgServer) ApproveOperation(ctx sdk.Context, op types.Operation) error {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		transfer, _ := pkg.GetTransfer(op)
		if err := k.ApproveTransferOperation(ctx, transfer); err != nil {
			return err
		}

		op.Status = types.OpStatus_APPROVED
		k.SetOperation(ctx, op)
	default:
		// Nothing to do
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationApproved,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
	))

	return nil
}

func (k msgServer) ApproveTransferOperation(ctx sdk.Context, transfer *types.Transfer) error {
	data, ok := k.tm.GetCollectionData(ctx, &tokentypes.CollectionDataIndex{Chain: transfer.From.Chain, Address: transfer.From.Address})
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection data does not exists")
	}

	from, ok := k.tm.GetOnChainItem(ctx, transfer.From)
	if !ok {
		// Item also does not exist
		item := tokentypes.Item{
			Index:      transfer.Origin,
			Collection: data.Collection,
			Meta:       transfer.Meta,
			OnChain:    []*tokentypes.OnChainItemIndex{transfer.From},
		}

		// Indexing seed and check if already exists. Meta not-nil validated during op creation
		if item.Meta.Seed != "" {
			if _, ok := k.tm.GetSeed(ctx, item.Meta.Seed); ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "seed already exists")
			}

			k.tm.SetSeed(ctx, tokentypes.Seed{
				Seed: item.Meta.Seed,
				Item: item.Index,
			})
			ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeSeedCreated,
				sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, item.Index),
				sdk.NewAttribute(tokentypes.AttributeKeySeed, item.Meta.Seed),
			))
		}

		k.tm.SetItem(ctx, item)
		ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeItemCreated,
			sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, item.Index),
		))

		from = tokentypes.OnChainItem{
			Index: transfer.From,
			Item:  item.Index,
		}

		k.tm.SetOnChainItem(ctx, from)
		ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeOnChainItemCreated,
			sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, from.Item),
			sdk.NewAttribute(tokentypes.AttributeKeyOnChainItemChain, from.Index.Chain),
		))
	}

	to, ok := k.tm.GetOnChainItem(ctx, transfer.From)
	if !ok {
		item, ok := k.tm.GetItem(ctx, from.Item)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "item does not exists but 'from' on chain item found")
		}

		item.OnChain = append(item.OnChain, transfer.To)
		k.tm.SetItem(ctx, item)

		to = tokentypes.OnChainItem{
			Index: transfer.To,
			Item:  item.Index,
		}

		k.tm.SetOnChainItem(ctx, to)
		ctx.EventManager().EmitEvent(sdk.NewEvent(tokentypes.EventTypeOnChainItemCreated,
			sdk.NewAttribute(tokentypes.AttributeKeyItemIndex, to.Item),
			sdk.NewAttribute(tokentypes.AttributeKeyOnChainItemChain, to.Index.Chain),
		))
	}

	return nil
}
