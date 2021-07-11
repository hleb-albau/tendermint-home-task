package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hleb-albau/registry/x/registry/types"
	"github.com/stretchr/testify/assert"
)

func createNChain(keeper *Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendChain(ctx, items[i])
	}
	return items
}

func TestChainGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetChain(ctx, item.Id))
	}
}

func TestChainExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasChain(ctx, item.Id))
	}
}

func TestChainRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.Id)
		assert.False(t, keeper.HasChain(ctx, item.Id))
	}
}

func TestChainGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllChain(ctx))
}

func TestChainCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetChainCount(ctx))
}
