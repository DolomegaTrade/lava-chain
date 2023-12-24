package types

import (
	"cosmossdk.io/math"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/lavanet/lava/x/downtime/v1"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	timerstoretypes "github.com/lavanet/lava/x/timerstore/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToModule(ctx sdk.Context, senderPool, recipientPool string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

type SpecKeeper interface {
	GetAllChainIDs(ctx sdk.Context) (chainIDs []string)
	GetSpec(ctx sdk.Context, index string) (val spectypes.Spec, found bool)
}

type TimerStoreKeeper interface {
	NewTimerStoreBeginBlock(storeKey storetypes.StoreKey, prefix string) *timerstoretypes.TimerStore
	NewTimerStoreEndBlock(storeKey storetypes.StoreKey, prefix string) *timerstoretypes.TimerStore
}

type EpochstorageKeeper interface {
	GetStakeStorageCurrent(ctx sdk.Context, chainID string) (epochstoragetypes.StakeStorage, bool)
}

type DowntimeKeeper interface {
	GetParams(ctx sdk.Context) (params v1.Params)
}

type StakingKeeper interface {
	BondedRatio(ctx sdk.Context) math.LegacyDec
	BondDenom(ctx sdk.Context) string
	// Methods imported from bank should be defined here
}

type DualStakingKeeper interface {
	RewardProvidersAndDelegators(ctx sdk.Context, providerAddr sdk.AccAddress, chainID string, totalReward math.Int, senderModule string, calcOnlyProvider bool, calcOnlyDelegators bool, calcOnlyContributer bool) (providerReward math.Int, totalRewards math.Int, err error)
	// Methods imported from bank should be defined here
}