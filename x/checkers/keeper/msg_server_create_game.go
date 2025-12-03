package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}
	if _, err := k.addressCodec.StringToBytes(msg.Black); err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidBlack, "invalid black address: %s", err)
	}
	if _, err := k.addressCodec.StringToBytes(msg.Red); err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidRed, "invalid red address: %s", err)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	systemInfo, err := k.SystemInfo.Get(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to get system info")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	newGame := rules.New()
	storedGame := types.StoredGame{
		Index: newIndex,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
	}

	err = storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.StoredGame.Set(ctx, storedGame.Index, storedGame)
	systemInfo.NextId++
	err = k.SystemInfo.Set(ctx, systemInfo)

	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to update system info")
	}

	_ = ctx

	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
