package keeper_test

import (
	"context"
	"testing"

	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t *testing.T) (types.MsgServer, context.Context) {
	f := initFixture(t)
	return keeper.NewMsgServerImpl(f.keeper), f.ctx
}

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
