package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListStoredGame(ctx context.Context, req *types.QueryAllStoredGameRequest) (*types.QueryAllStoredGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storedGames, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.StoredGame,
		req.Pagination,
		func(_ string, value types.StoredGame) (types.StoredGame, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredGameResponse{StoredGame: storedGames, Pagination: pageRes}, nil
}

func (q queryServer) GetStoredGame(ctx context.Context, req *types.QueryGetStoredGameRequest) (*types.QueryGetStoredGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.StoredGame.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetStoredGameResponse{StoredGame: val}, nil
}
