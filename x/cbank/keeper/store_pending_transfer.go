package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/cbank/types"
)

func (k Keeper) SetPendingTransfer(ctx sdk.Context, pt *types.PendingTransfer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingTransferKeyPrefix))
	b := k.cdc.MustMarshal(pt)
	store.Set(types.PendingTransferKey(pt.Index), b)
}

func (k Keeper) GetPendingTransfer(ctx sdk.Context, index string) (val types.PendingTransfer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingTransferKeyPrefix))

	b := store.Get(types.PendingTransferKey(index))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemovePendingTransfer(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingTransferKeyPrefix))
	store.Delete(types.PendingTransferKey(index))
}

func (k Keeper) GetAllPendingTransfer(ctx sdk.Context) (list []types.PendingTransfer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingTransferKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingTransfer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
