package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/alice/checkers/x/checkers/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	if genState.SystemInfo.NextId != 0 {
		if err := k.SystemInfo.Set(ctx, genState.SystemInfo); err != nil {
			return err
		}
	}
	for _, elem := range genState.StoredGameMap {
		if err := k.StoredGame.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	systemInfo, err := k.SystemInfo.Get(ctx)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return nil, err
	}
	genesis.SystemInfo = systemInfo
	if err := k.StoredGame.Walk(ctx, nil, func(_ string, val types.StoredGame) (stop bool, err error) {
		genesis.StoredGameMap = append(genesis.StoredGameMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
