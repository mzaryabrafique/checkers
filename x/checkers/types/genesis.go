package types

import "fmt"

// DefaultGenesis returns the default genesis state
const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		SystemInfo: SystemInfo{
			NextId: uint64(DefaultIndex),
		},
		StoredGameMap: []StoredGame{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	storedGameIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredGameMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := storedGameIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedGame")
		}
		storedGameIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
