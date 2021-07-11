package cli_test

import (
	"fmt"
	utils "github.com/hleb-albau/registry/testutil"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/hleb-albau/registry/testutil/network"
	"github.com/hleb-albau/registry/x/registry/client/cli"
)

func TestRegisterChain(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	for _, tc := range []struct {
		desc string
		args []string
		err  error
		code uint32
	}{
		{
			desc: "valid",
			args: commonArgs(net),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := append([]string{"someChainID"}, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRegisterChain(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdateChain(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	chainID := "someChainID"
	common := commonArgs(net)
	args := append([]string{chainID}, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRegisterChain(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc    string
		chainID string
		args    []string
		code    uint32
		err     error
	}{
		{
			desc:    "valid",
			chainID: chainID,
			args:    common,
		},
		{
			desc:    "key not found",
			chainID: "1",
			args:    common,
			code:    sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.chainID}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateChain(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

// could be added tests all code paths, impl only positive one
func TestTransferChainOwnership(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	chainID := "someChainID"
	common := commonArgs(net)
	args := append([]string{chainID}, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRegisterChain(), args)
	require.NoError(t, err)

	args = append([]string{chainID, utils.NewAccAddress()}, common...)
	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdTransferChainOwnership(), args)
	require.NoError(t, err)
	var resp sdk.TxResponse
	require.NoError(t, ctx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &resp))
	require.Equal(t, uint32(0), resp.Code)
}

func commonArgs(net *network.Network) []string {
	return []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, net.Validators[0].Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
}
