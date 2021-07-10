# Chain Registry Module Draft

The goal of this module is to be a simple registry for chains with the ability to transfer ownership.

# State

## Chain

`Chain` contains various information about a given chain. `ChainID` used to uniquely identify a chain.

Key: `chain-[ChainID] -> Chain`

```go
type Chain struct {
ChainID string
Owner string
// other properties, required by final module impl
}
```

# Messages

## MsgRegisterChain

Register a new chain with given **unoccupied** `ChainID`(there is no already registered chain with the same `ChainID`).

```go
type MsgRegisterChain struct {
ChainID string
Owner string
// other fields, required by final module impl
}
```

## MsgUpdateChain

Updates information for a given chain. Note, `ChainID` and `Owner` is **unchangeable** fields. Invocation permitted only
to 'Chain.Owner' account.

```go
type MsgUpdateChain struct {
ChainID string
Owner string
// other updatable fields, required by final module impl
}
```

## MsgTransferChainOwnership

Transfer ownership for a chain identified by `ChainID` to `NewOwner` account. Invocation permitted only to 'Chain.Owner'
account.

```go
type MsgTransferChainOwnership struct {
ChainID string
Owner string
NewOwner string
}
```

# Ideas

* The initial price of chain registration is inversely proportional to `ChainID` length. All deal tokens go to the
  community pool.
* Any chain transferring could be charged with some additional fee in favor of community pool. Fee amount (or
  percentage) should be chain parameter, thus can be changed via government procedures.
* `Chain` itself can be [NFT token](https://github.com/cosmos/cosmos-sdk/pull/9329) (gives future support for ibc)
* Register `ChainID` as alias to existing `Chain`. Convert existing `Chain` to be alias to another `Chain`.
* Fund a `Chain` with ability to claim tokens by owner.
* Partial ownership of a chain. Each owner has a share, and can claim tokens up to her % of share?
