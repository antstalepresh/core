package wasmbinding

import (
	denomkeeper "github.com/Team-Kujira/core/x/denom/keeper"

	oraclekeeper "github.com/Team-Kujira/core/x/oracle/keeper"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func RegisterCustomPlugins(
	bank bankkeeper.BaseKeeper,
	oracle oraclekeeper.Keeper,
	denom denomkeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(bank, oracle, denom)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})

	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, denom),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
