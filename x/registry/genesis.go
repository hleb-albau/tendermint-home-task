package registry

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hleb-albau/registry/x/registry/keeper"
	"github.com/hleb-albau/registry/x/registry/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the chain
	for _, elem := range genState.ChainList {
		k.SetChain(ctx, *elem)
	}

	// Set chain count
	k.SetChainCount(ctx, genState.ChainCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all chain
	chainList := k.GetAllChain(ctx)
	for _, elem := range chainList {
		elem := elem
		genesis.ChainList = append(genesis.ChainList, &elem)
	}

	// Set the current count
	genesis.ChainCount = k.GetChainCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
