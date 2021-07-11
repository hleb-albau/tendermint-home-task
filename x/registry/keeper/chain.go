package keeper

import (
	"encoding/binary"
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

// AppendChain appends a chain in the store with a new id and update the count
func (k Keeper) AppendChain(
	ctx sdk.Context,
	chain types.Chain,
) uint64 {
	// Create the chain
	count := k.GetChainCount(ctx)

	// Set the ID of the appended value
	chain.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(GetChainIDBytes(chain.Id), appendedValue)

	// Update chain count
	k.SetChainCount(ctx, count+1)

	return count
}

// SetChain set a specific chain in the store
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	b := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(GetChainIDBytes(chain.Id), b)
}

// GetChain returns a chain from its id
func (k Keeper) GetChain(ctx sdk.Context, id uint64) types.Chain {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	var chain types.Chain
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetChainIDBytes(id)), &chain)
	return chain
}

// HasChain checks if the chain exists in the store
func (k Keeper) HasChain(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	return store.Has(GetChainIDBytes(id))
}

// GetChainOwner returns the creator of the chain
func (k Keeper) GetChainOwner(ctx sdk.Context, id uint64) string {
	return k.GetChain(ctx, id).Creator
}

// RemoveChain removes a chain from the store
func (k Keeper) RemoveChain(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	store.Delete(GetChainIDBytes(id))
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

// GetChainIDBytes returns the byte representation of the ID
func GetChainIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetChainIDFromBytes returns ID in uint64 format from a byte array
func GetChainIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
