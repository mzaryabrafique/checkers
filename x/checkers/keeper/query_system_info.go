package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"

	"github.com/alice/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) GetSystemInfo(ctx context.Context, req *types.QueryGetSystemInfoRequest) (*types.QueryGetSystemInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.SystemInfo.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetSystemInfoResponse{SystemInfo: val}, nil
}
