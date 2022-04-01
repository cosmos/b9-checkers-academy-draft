package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	"github.com/xavierlepretre/checkers/app"
	"github.com/xavierlepretre/checkers/x/checkers/keeper"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func ModuleAccountAddrs(maccPerms map[string][]string) map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func setupKeeper2(t testing.TB, bankGenesis *banktypes.GenesisState) (*keeper.Keeper, sdk.Context) {
	// module account permissions
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
		types.ModuleName:           nil,
	}

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	registry := codectypes.NewInterfaceRegistry()
	marshaller := codec.NewProtoCodec(registry)
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	// Params
	paramsStoreKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	paramsMemStoreKey := storetypes.NewMemoryStoreKey(paramstypes.TStoreKey)
	stateStore.MountStoreWithDB(paramsStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(paramsMemStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())
	paramsKeeper := paramskeeper.NewKeeper(marshaller, encoding.Amino, paramsStoreKey, paramsMemStoreKey)
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(types.ModuleName)

	// Auth
	authStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	authMemStoreKey := storetypes.NewMemoryStoreKey("transient_auth")
	stateStore.MountStoreWithDB(authStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(authMemStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())
	authSubSpace, ok := paramsKeeper.GetSubspace(authtypes.ModuleName)
	require.True(t, ok, "Could not get auth subspace")
	accountKeeper := authkeeper.NewAccountKeeper(
		marshaller, authStoreKey, authSubSpace, authtypes.ProtoBaseAccount, maccPerms,
	)

	// Bank
	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(bankStoreKey, sdk.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())
	bankSubSpace, ok := paramsKeeper.GetSubspace(banktypes.ModuleName)
	require.True(t, ok, "Could not get bank subspace")
	bankKeeper := bankkeeper.NewBaseKeeper(
		marshaller, bankStoreKey, accountKeeper, bankSubSpace, ModuleAccountAddrs(maccPerms),
	)

	// Checkers
	checkersStoreKey := sdk.NewKVStoreKey(types.StoreKey)
	checkersMemStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	stateStore.MountStoreWithDB(checkersStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(checkersMemStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	keeper := keeper.NewKeeper(
		bankKeeper,
		marshaller,
		checkersStoreKey,
		checkersMemStoreKey,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	bankKeeper.InitGenesis(ctx, bankGenesis)
	return keeper, ctx
}

func setupKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	keeper := keeper.NewKeeper(
		*new(bankkeeper.Keeper),
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return keeper, ctx
}
