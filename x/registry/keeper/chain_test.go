package keeper

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hleb-albau/registry/x/registry/types"
	"github.com/stretchr/testify/assert"
)

func createNChain(keeper *Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].Owner = NewAccAddress()
		items[i].ChainID = strconv.Itoa(i)
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func TestChainGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetChain(ctx, item.ChainID))
	}
}

func TestChainExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasChain(ctx, item.ChainID))
	}
}

func TestChainRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.ChainID)
		assert.False(t, keeper.HasChain(ctx, item.ChainID))
	}
}

func TestChainGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllChain(ctx))
}

func NewAccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}
