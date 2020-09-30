package app

import (
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

func loadKVStoreKeys(keys ...string) map[string]*sdk.KVStoreKey {
	defaultKeys := []string{
		bam.MainStoreKey,
		auth.StoreKey,
		staking.StoreKey,
		supply.StoreKey,
		params.StoreKey,
	}
	for _, key := range keys {
		defaultKeys = append(defaultKeys, key)
	}
	return sdk.NewKVStoreKeys(defaultKeys...)
}

func loadModules(modules ...module.AppModuleBasic) module.BasicManager {
	// These are defaults
	moduleImports := []module.AppModuleBasic{
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		params.AppModuleBasic{},
		supply.AppModuleBasic{},
	}
	for _, m := range modules {
		moduleImports = append(moduleImports, m)
	}

	return module.NewBasicManager(moduleImports...)
}

func (app *NewApp) loadCustomManagers(managers ...module.AppModule) {
	// These are defaults
	managerImports := []module.AppModule{
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),
	}
	for _, m := range managers {
		managerImports = append(managerImports, m)
	}

	app.mm = module.NewManager(managerImports...)
}

// func (app *NewApp) loadDefaultKeepers(keys map[string]*sdk.KVStoreKey, tKeys map[string]*sdk.TransientStoreKey) {
func (app *NewApp) loadDefaultKeepers() {
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keys[params.StoreKey], app.tKeys[params.TStoreKey])
	app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.paramsKeeper.Subspace(staking.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount,
	)

	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.subspaces[bank.ModuleName],
		app.ModuleAccountAddrs(),
	)

	app.supplyKeeper = supply.NewKeeper(
		app.cdc,
		app.keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		maccPerms,
	)

	stakingKeeper := staking.NewKeeper(
		app.cdc,
		app.keys[staking.StoreKey],
		app.supplyKeeper,
		app.subspaces[staking.ModuleName],
	)

	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(),
	)
}
