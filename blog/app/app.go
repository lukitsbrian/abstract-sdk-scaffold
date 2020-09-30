package app

import (
	"io"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog"
	blogkeeper "github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/keeper"
	blogtypes "github.com/lukitsbrian/abstract-sdk-scaffold/blog/x/blog/types"
	// this line is used by starport scaffolding # 1
)

const appName = "blog"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.blogcli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.blogd")
	ModuleBasics    = loadModules(
		blog.AppModuleBasic{},
		// this line is used by starport scaffolding # 2
	)

	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
	}
)

type MyApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	keys  map[string]*sdk.KVStoreKey
	tKeys map[string]*sdk.TransientStoreKey

	subspaces map[string]params.Subspace

	accountKeeper auth.AccountKeeper
	bankKeeper    bank.Keeper
	stakingKeeper staking.Keeper
	supplyKeeper  supply.Keeper
	paramsKeeper  params.Keeper
	blogKeeper    blogkeeper.Keeper
	// this line is used by starport scaffolding # 3
	mm *module.Manager

	sm *module.SimulationManager
}

var _ simapp.App = (*MyApp)(nil)

func NewInitApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) *MyApp {

	customKVStoreKeys := []string{
		blogtypes.StoreKey,
	}

	customTStoreKeys := []string{}

	// Create app
	var app = &MyApp{}

	app.Init(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	app.loadKeys(customKVStoreKeys, customTStoreKeys)

	// load keepers
	app.loadDefaultKeepers()

	app.blogKeeper = blogkeeper.NewKeeper(
		app.bankKeeper,
		app.cdc,
		app.keys[blogtypes.StoreKey],
	)

	// this line is used by starport scaffolding # 4

	// load modules
	app.loadCustomManagers(
		blog.NewAppModule(app.blogKeeper, app.bankKeeper),
		// this line is used by starport scaffolding # 6
	)

	app.mm.SetOrderEndBlockers(staking.ModuleName)

	app.mm.SetOrderInitGenesis(
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		blogtypes.ModuleName,
		supply.ModuleName,
		genutil.ModuleName,
		// this line is used by starport scaffolding # 7
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.Setup()

	app.Mount()

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

func (app *MyApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState

	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *MyApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *MyApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *MyApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

func (app *MyApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (app *MyApp) Codec() *codec.Codec {
	return app.cdc
}

func (app *MyApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

func GetMaccPerms() map[string][]string {
	modAccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		modAccPerms[k] = v
	}
	return modAccPerms
}

func (app *MyApp) Init(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp)) {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	app.BaseApp = bApp
	app.cdc = cdc
	app.invCheckPeriod = invCheckPeriod
	app.subspaces = make(map[string]params.Subspace)
}

func (app *MyApp) Setup() {
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)
}

func (app *MyApp) Mount() {
	app.MountKVStores(app.keys)
	app.MountTransientStores(app.tKeys)
}
