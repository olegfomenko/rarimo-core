package cbank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cbank/keeper"
	"github.com/rarimo/rarimo-core/x/cbank/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, com := range genState.Commitments {
		k.SetCommitment(ctx, &com)
	}
	for _, pt := range genState.Transfers {
		k.SetPendingTransfer(ctx, &pt)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Commitments = k.GetAllCommitment(ctx)
	genesis.Transfers = k.GetAllPendingTransfer(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
