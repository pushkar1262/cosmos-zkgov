package zkgov

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/zkgov/keeper"
)

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
}

type ModuleOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
	)

	return ModuleOutputs{
		Keeper: k,
		Module: NewAppModule(in.Cdc, k),
	}
}

// Note: module registration via app wiring is intentionally omitted here to avoid
// requiring a protobuf config type with cosmos.app.v1alpha1.module option.
