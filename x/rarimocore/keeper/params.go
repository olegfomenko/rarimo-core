package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) UpdateParams(ctx sdk.Context, params types.Params) {
	currentParams := k.GetParams(ctx)
	activePartiesAmount := getActivePartiesAmount(params.Parties)

	params.Threshold = uint64(crypto.GetThreshold(activePartiesAmount))
	params.IsUpdateRequired = activePartiesAmount != getActivePartiesAmount(currentParams.Parties)

	k.SetParams(ctx, params)
}

func (k Keeper) GetKeyECDSA(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.GetKeyECDSA()
}

func (k Keeper) UpdateLastSignature(ctx sdk.Context, sig string) {
	params := k.GetParams(ctx)
	params.LastSignature = sig
	k.SetParams(ctx, params)
}

func (k Keeper) getActivePartiesAmount(ctx sdk.Context) int {
	params := k.GetParams(ctx)
	return getActivePartiesAmount(params.Parties)
}

func getActivePartiesAmount(parties []*types.Party) int {
	activePartyCount := 0

	for _, party := range parties {
		if party.Status == types.PartyStatus_Active {
			activePartyCount++
		}
	}

	return activePartyCount
}
