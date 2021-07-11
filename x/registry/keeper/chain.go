package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hleb-albau/registry/x/registry/types"
	"strconv"
)

// GetChainCount get the total number of chain
func (k Keeper) GetChainCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainCountKey))
	byteKey := types.KeyPrefix(types.ChainCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetChainCount set the total number of chain
func (k Keeper) SetChainCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainCountKey))
	byteKey := types.KeyPrefix(types.ChainCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// SetChain set a specific chain in the store
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	b := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(GetChainKey(chain.ChainID), b)
}

// GetChain returns a chain from its id
func (k Keeper) GetChain(ctx sdk.Context, chainID string) types.Chain {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	var chain types.Chain
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetChainKey(chainID)), &chain)
	return chain
}

// HasChain checks if the chain exists in the store
func (k Keeper) HasChain(ctx sdk.Context, chainID string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	return store.Has(GetChainKey(chainID))
}

// GetChainOwner returns the creator of the chain
func (k Keeper) GetChainOwner(ctx sdk.Context, chainID string) string {
	return k.GetChain(ctx, chainID).Owner
}

// RemoveChain removes a chain from the store
func (k Keeper) RemoveChain(ctx sdk.Context, chainID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	store.Delete(GetChainKey(chainID))
}

// GetAllChain returns all chain
func (k Keeper) GetAllChain(ctx sdk.Context) (list []types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Chain
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetChainKey(chainID string) []byte {
	return []byte(chainID)
}
