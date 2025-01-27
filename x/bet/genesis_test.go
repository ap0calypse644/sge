package bet_test

import (
	"testing"

	"github.com/sge-network/sge/testutil/nullify"
	simappUtil "github.com/sge-network/sge/testutil/simapp"
	"github.com/sge-network/sge/x/bet"
	"github.com/sge-network/sge/x/bet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BetList: []types.Bet{
			{
				UID: "0",
			},
			{
				UID: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	bet.InitGenesis(ctx, tApp.BetKeeper, genesisState)
	got := bet.ExportGenesis(ctx, tApp.BetKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BetList, got.BetList)
	// this line is used by starport scaffolding # genesis/test/assert
}
