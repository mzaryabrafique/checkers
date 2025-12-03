package types

import "cosmossdk.io/collections"

// StoredGameKey is the prefix to retrieve all StoredGame
var StoredGameKey = collections.NewPrefix("storedGame/value/")
