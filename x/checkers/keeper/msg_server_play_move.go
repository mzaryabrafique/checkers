package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/alice/checkers/x/checkers/types"
)

func (k msgServer) PlayMove(ctx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgPlayMoveResponse{}, nil
}
