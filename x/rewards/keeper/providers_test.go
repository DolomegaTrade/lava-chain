package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/lavanet/lava/testutil/common"
	"github.com/lavanet/lava/x/rewards/types"
	subscription "github.com/lavanet/lava/x/subscription/keeper"
	"github.com/stretchr/testify/require"
)

// for this test there are no relays, this means no rewards will be given to the providers, and this means no bonus rewards should be sent
func TestZeroProvidersRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment from the subscription
	// example: chain starts at block 1 and next distribution is in 100blocks, block 101 (1 month)
	// a user bought a subscription in block 2, it will expire in block 102
	// in block 101 we will distribute rewards but since there hasnt been any subscription monthly expiration no rewards were sent and thus the first month is always 0 rewards
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)

	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)
}

// the rewards here is maxboost*totalbaserewards, in this test the rewards for the providers are low (first third of the graph)
func TestBasicBoostProvidersRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	baserewards := uint64(100)
	// the rewards by the subscription will be limited by LIMIT_TOKEN_PER_CU
	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, baserewards)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	expectedReward, _, _ := ts.DeductParticipationFees(sdk.NewIntFromUint64(baserewards * subscription.LIMIT_TOKEN_PER_CU))
	require.Equal(t, expectedReward, res.Rewards[0].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount.Amount.Uint64(), baserewards*subscription.LIMIT_TOKEN_PER_CU*ts.Keepers.Rewards.GetParams(ts.Ctx).MaxRewardBoost)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// the rewards here is spec payout allocation (full rewards from the pool), in this test the rewards for the providers are medium (second third of the graph)
func TestSpecAllocationProvidersRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	expectedReward, _, _ := ts.DeductParticipationFees(ts.plan.Price.Amount)
	require.Equal(t, expectedReward, res.Rewards[0].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, res.Rewards[0].Amount.Amount, distBalance)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// the rewards here is the diminishing part of the reward, in this test the rewards for the providers are high (third third of the graph)
func TestProvidersDiminishingRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	for i := 0; i < 7; i++ {
		consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
		_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
		require.Nil(t, err)

		msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
		_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
		require.Nil(t, err)
	}

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)

	expectedReward, _, _ := ts.DeductParticipationFees(ts.plan.Price.Amount)
	expectedReward = expectedReward.MulRaw(7) // the participation fees are done separately on each of the 7 relays
	require.Equal(t, expectedReward, res.Rewards[0].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)

	require.Equal(t, res.Rewards[0].Amount.Amount, sdk.NewDecWithPrec(15, 1).MulInt(distBalance).Sub(sdk.NewDecWithPrec(5, 1).MulInt(ts.plan.Price.Amount.MulRaw(7))).TruncateInt())
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// the rewards here is the zero since base rewards are very big, in this test the rewards for the providers are at the end of the graph
func TestProvidersEndRewards(t *testing.T) {
	ts := newTester(t)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	for i := 0; i < 50; i++ {
		consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
		_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
		require.Nil(t, err)

		msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
		_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
		require.Nil(t, err)
	}

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	expectedReward, _, _ := ts.DeductParticipationFees(ts.plan.Price.Amount)
	expectedReward = expectedReward.MulRaw(50) // the participation fees are done separately on each of the 50 relays
	require.Equal(t, expectedReward, res.Rewards[0].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// in this test we create 2 specs with 1 provider each, one of the specs shares is zero
// this means that no matter how much rewards the providers in this spec will get, they will get 0 bonus rewards
func Test2SpecsZeroShares(t *testing.T) {
	ts := newTester(t)
	spec2 := ts.spec
	spec2.Index = "mock2"
	spec2.Name = spec2.Index
	spec2.Shares = 0
	ts.AddSpec(spec2.Index, spec2)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, 2*testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	err = ts.StakeProvider(providerAcc.Addr.String(), spec2, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	consumerAcc2, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc2.Addr.String(), consumerAcc2.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc.Addr.String(), consumerAcc2, []string{spec2.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 2)
	expectedReward, _, _ := ts.DeductParticipationFees(ts.plan.Price.Amount)
	require.Equal(t, expectedReward, res.Rewards[0].Amount.Amount)
	require.Equal(t, expectedReward, res.Rewards[1].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	require.Equal(t, distBalance, res.Rewards[0].Amount.Amount)
	require.Equal(t, res.Rewards[0].ChainId, ts.spec.Index)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// 2 specs with one of them double the shares than the other
// the providers will have the same amount of CU used, thus the same rewards
// the bonus for the provider with double the shares should be double than the other provider
func Test2SpecsDoubleShares(t *testing.T) {
	ts := newTester(t)
	spec2 := ts.spec
	spec2.Index = "mock2"
	spec2.Name = spec2.Index
	spec2.Shares *= 2
	ts.AddSpec(spec2.Index, spec2)

	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, 2*testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	err = ts.StakeProvider(providerAcc.Addr.String(), spec2, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	consumerAcc2, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc2.Addr.String(), consumerAcc2.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc.Addr.String(), consumerAcc2, []string{spec2.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 2)
	expectedReward, _, _ := ts.DeductParticipationFees(ts.plan.Price.Amount)
	require.Equal(t, expectedReward, res.Rewards[0].Amount.Amount)
	require.Equal(t, expectedReward, res.Rewards[1].Amount.Amount)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), "")
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 2)
	require.Equal(t, res.Rewards[0].Amount.Amount, res.Rewards[1].Amount.Amount.MulRaw(2))
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)
}

// in this test we setup 3 providers, each with different cu used (-> 3 different rewards from the plan) (1,2,4)
// the providers should get bonus rewards according to their plan rewards
func TestBonusRewards3Providers(t *testing.T) {
	ts := newTester(t)

	providerAcc1, _ := ts.AddAccount(common.PROVIDER, 1, 2*testBalance)
	err := ts.StakeProvider(providerAcc1.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	providerAcc2, _ := ts.AddAccount(common.PROVIDER, 2, 2*testBalance)
	err = ts.StakeProvider(providerAcc2.Addr.String(), ts.spec, 2*testBalance)
	require.Nil(t, err)

	providerAcc3, _ := ts.AddAccount(common.PROVIDER, 3, 3*testBalance)
	err = ts.StakeProvider(providerAcc3.Addr.String(), ts.spec, 3*testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	msg := ts.SendRelay(providerAcc1.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()/2)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc2.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64())
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	msg = ts.SendRelay(providerAcc3.Addr.String(), consumerAcc, []string{ts.spec.Index}, ts.plan.Price.Amount.Uint64()*2)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc1.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc1.Addr.String(), "")
	require.Nil(t, err)

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc2.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc2.Addr.String(), "")
	require.Nil(t, err)

	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc3.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc3.Addr.String(), "")
	require.Nil(t, err)

	// now the provider should get all of the provider allocation
	ts.AdvanceMonths(1)
	distBalance := ts.Keepers.Rewards.TotalPoolTokens(ts.Ctx, types.ProviderRewardsDistributionPool)
	ts.AdvanceEpoch()

	res1, err := ts.QueryDualstakingDelegatorRewards(providerAcc1.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	// we sub 3 because of truncating
	require.Equal(t, res1.Rewards[0].Amount.Amount, distBalance.QuoRaw(7).SubRaw(3))
	_, err = ts.TxDualstakingClaimRewards(providerAcc1.Addr.String(), providerAcc1.Addr.String())
	require.Nil(t, err)

	res2, err := ts.QueryDualstakingDelegatorRewards(providerAcc2.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	// we sub 1 because of truncating
	require.Equal(t, res2.Rewards[0].Amount.Amount, distBalance.QuoRaw(7).MulRaw(2))
	_, err = ts.TxDualstakingClaimRewards(providerAcc2.Addr.String(), providerAcc2.Addr.String())
	require.Nil(t, err)

	res3, err := ts.QueryDualstakingDelegatorRewards(providerAcc3.Addr.String(), "", "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	// we add 6 because of truncating
	require.Equal(t, res3.Rewards[0].Amount.Amount, distBalance.QuoRaw(7).MulRaw(4).AddRaw(6))
	_, err = ts.TxDualstakingClaimRewards(providerAcc3.Addr.String(), providerAcc3.Addr.String())
	require.Nil(t, err)
}

// TestValidatorsAndCommunityParticipation checks that the validators and community participation funds
// are as expected (according to communityTax and validatorsSubscriptionParticipation params)
func TestValidatorsAndCommunityParticipation(t *testing.T) {
	ts := newTester(t)

	// set the communityTax and validatorsSubscriptionParticipation params to const values
	// communityTax = 50%
	// validatorsSubscriptionParticipation = 10%
	distParams := distributiontypes.DefaultParams()
	distParams.CommunityTax = sdk.NewDecWithPrec(5, 1) // 0.5
	err := ts.Keepers.Distribution.SetParams(ts.Ctx, distParams)
	require.Nil(t, err)

	paramKey := string(types.KeyValidatorsSubscriptionParticipation)
	newDecParam, err := sdk.NewDecWithPrec(1, 1).MarshalJSON() // 0.1
	require.Nil(ts.T, err)
	paramVal := string(newDecParam)
	err = ts.TxProposalChangeParam(types.ModuleName, paramKey, paramVal)
	require.Nil(ts.T, err)

	// create provider+comsumer, send relay and send relay payment TX
	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err = ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64())
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, false)
	require.Nil(t, err)

	baserewards := uint64(100)
	// the rewards by the subscription will be limited by LIMIT_TOKEN_PER_CU
	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, baserewards)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment from the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	expectedReward := sdk.NewIntFromUint64(baserewards * subscription.LIMIT_TOKEN_PER_CU)
	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	deductedReward, validatorsParticipation, communityParticipation := ts.DeductParticipationFees(expectedReward)
	require.True(t, expectedReward.Equal(deductedReward.Add(validatorsParticipation).Add(communityParticipation)))

	// check participation percentages values accoding to hard-coded values
	// validators participation percentage = (validatorsSubscriptionParticipation / (1 - communityTax)) = (10% / 100% - 50%) = 0.2
	// community participation percentage = (validatorsSubscriptionParticipation + communityTax) - validators participation percentage = (10% + 50%) - 0.2 = 0.4
	validatorsPerc := validatorsParticipation.MulRaw(100).Quo(expectedReward)
	communityPerc := communityParticipation.MulRaw(100).Quo(expectedReward)
	require.Equal(t, int64(20), validatorsPerc.Int64())
	require.Equal(t, int64(40), communityPerc.Int64())

	// check actual balance of the commuinty pool
	// community pool should have 40% of expected reward
	communityCoins := ts.Keepers.Distribution.GetFeePoolCommunityCoins(ts.Ctx)
	communityBalance := communityCoins.AmountOf(ts.TokenDenom()).TruncateInt()
	require.True(t, expectedReward.Mul(communityPerc).QuoRaw(100).Equal(communityBalance))
}

func TestBonusReward49months(t *testing.T) {
	ts := newTester(t)
	providerAcc, _ := ts.AddAccount(common.PROVIDER, 1, testBalance)
	err := ts.StakeProvider(providerAcc.Addr.String(), ts.spec, testBalance)
	require.Nil(t, err)

	ts.AdvanceEpoch()

	consumerAcc, _ := ts.AddAccount(common.CONSUMER, 1, ts.plan.Price.Amount.Int64()*100)
	_, err = ts.TxSubscriptionBuy(consumerAcc.Addr.String(), consumerAcc.Addr.String(), ts.plan.Index, 1, true)
	require.Nil(t, err)

	for i := 0; i < 50; i++ {
		ts.AdvanceMonths(1)
		ts.AdvanceBlocks(ts.BlocksToSave() + 1)
	}

	baserewards := uint64(100)
	// the rewards by the subscription will be limited by LIMIT_TOKEN_PER_CU
	msg := ts.SendRelay(providerAcc.Addr.String(), consumerAcc, []string{ts.spec.Index}, baserewards)
	_, err = ts.TxPairingRelayPayment(msg.Creator, msg.Relays...)
	require.Nil(t, err)

	// first months there are no bonus rewards, just payment ffrom the subscription
	ts.AdvanceMonths(1)
	ts.AdvanceBlocks(ts.BlocksToSave() + 1)

	res, err := ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 1)
	_, err = ts.TxDualstakingClaimRewards(providerAcc.Addr.String(), providerAcc.Addr.String())
	require.Nil(t, err)

	// now the provider should get all of the provider allocation (but there arent any)
	ts.AdvanceMonths(1)
	ts.AdvanceEpoch()

	// there should be no bonus rewards
	res, err = ts.QueryDualstakingDelegatorRewards(providerAcc.Addr.String(), providerAcc.Addr.String(), "")
	require.Nil(t, err)
	require.Len(t, res.Rewards, 0)
}
