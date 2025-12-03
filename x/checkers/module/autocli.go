package checkers

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/alice/checkers/x/checkers/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "GetSystemInfo",
					Use:       "get-system-info",
					Short:     "Gets a systemInfo",
					Alias:     []string{"show-system-info"},
				},
				{
					RpcMethod: "ListStoredGame",
					Use:       "list-stored-game",
					Short:     "List all storedGame",
				},
				{
					RpcMethod:      "GetStoredGame",
					Use:            "get-stored-game [id]",
					Short:          "Gets a storedGame",
					Alias:          []string{"show-stored-game"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateGame",
					Use:            "create-game [black] [red]",
					Short:          "Send a createGame tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "black"}, {ProtoField: "red"}},
				},
				{
					RpcMethod:      "PlayMove",
					Use:            "play-move [game-index] [from-x] [from-y] [to-x] [to-y]",
					Short:          "Send a playMove tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "game_index"}, {ProtoField: "from_x"}, {ProtoField: "from_y"}, {ProtoField: "to_x"}, {ProtoField: "to_y"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
