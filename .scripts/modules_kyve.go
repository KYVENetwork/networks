package main

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmOs "github.com/tendermint/tendermint/libs/os"

	// Bundles
	bundlesTypes "github.com/KYVENetwork/chain/x/bundles/types"
	// Delegation
	delegationTypes "github.com/KYVENetwork/chain/x/delegation/types"
	// Global
	globalTypes "github.com/KYVENetwork/chain/x/global/types"
	// Pool
	poolTypes "github.com/KYVENetwork/chain/x/pool/types"
	// Stakers
	stakersTypes "github.com/KYVENetwork/chain/x/stakers/types"
	// Staking
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	// Team
	teamTypes "github.com/KYVENetwork/chain/x/team/types"
)

func GenerateBundlesState() []byte {
	bundlesState := bundlesTypes.DefaultGenesis()

	var rawBundlesState bytes.Buffer
	_ = marshaler.Marshal(&rawBundlesState, bundlesState)

	return rawBundlesState.Bytes()
}

func GenerateDelegationState() []byte {
	delegationState := delegationTypes.DefaultGenesis()

	delegationState.Params = delegationTypes.Params{
		UnbondingDelegationTime: 7 * 24 * 60 * 60,
		RedelegationCooldown:    7 * 24 * 60 * 60,
		RedelegationMaxAmount:   3,
		VoteSlash:               "0.02",
		UploadSlash:             "0.05",
		TimeoutSlash:            "0.0005",
	}

	var rawDelegationState bytes.Buffer
	_ = marshaler.Marshal(&rawDelegationState, delegationState)

	return rawDelegationState.Bytes()
}

func GenerateGlobalState() []byte {
	globalState := globalTypes.DefaultGenesisState()

	globalState.Params.MinGasPrice = sdk.MustNewDecFromStr("0.02")
	globalState.Params.GasAdjustments = []globalTypes.GasAdjustment{
		{
			Type:   sdk.MsgTypeURL(&stakingTypes.MsgCreateValidator{}),
			Amount: BlockMaxGas * 2,
		},
		{
			Type:   sdk.MsgTypeURL(&stakersTypes.MsgCreateStaker{}),
			Amount: 50_000_000,
		},
	}
	globalState.Params.MinInitialDepositRatio = sdk.MustNewDecFromStr("0.50")

	var rawGlobalState bytes.Buffer
	_ = marshaler.Marshal(&rawGlobalState, globalState)

	return rawGlobalState.Bytes()
}

func GeneratePoolState() []byte {
	poolState := poolTypes.DefaultGenesis()

	var rawPoolState bytes.Buffer
	_ = marshaler.Marshal(&rawPoolState, poolState)

	return rawPoolState.Bytes()
}

func GenerateStakersState() []byte {
	stakersState := stakersTypes.DefaultGenesis()

	stakersState.Params = stakersTypes.Params{
		CommissionChangeTime: 7 * 24 * 60 * 60,
		LeavePoolTime:        7 * 24 * 60 * 60,
	}

	var rawStakersState bytes.Buffer
	_ = marshaler.Marshal(&rawStakersState, stakersState)

	return rawStakersState.Bytes()
}

func GenerateTeamState(chainID string) []byte {
	teamState := teamTypes.DefaultGenesis()

	teamAccounts, err := InjectTeamAccounts(chainID)
	if err != nil {
		fmt.Println("❌ Failed to inject team accounts!")
		tmOs.Exit(err.Error())
	}

	teamState.AccountList = teamAccounts
	teamState.AccountCount = uint64(len(teamAccounts))

	var rawTeamState bytes.Buffer
	_ = marshaler.Marshal(&rawTeamState, teamState)

	return rawTeamState.Bytes()
}
