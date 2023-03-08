package main

import (
	"bytes"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmOs "github.com/tendermint/tendermint/libs/os"

	// Auth
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// Authz
	authzTypes "github.com/cosmos/cosmos-sdk/x/authz"
	// Bank
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// Capability
	capabilityTypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	// Crisis
	crisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	// Distribution
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	// Evidence
	evidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	// FeeGrant
	feeGrantTypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	// Gov
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	// Group
	groupTypes "github.com/cosmos/cosmos-sdk/x/group"
	// Mint
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	// Slashing
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	// Staking
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func GenerateAuthState(chainID string, denom string) []byte {
	authState := authTypes.DefaultGenesisState()

	accounts, err := InjectGenesisAccounts(chainID, denom)
	if err != nil {
		fmt.Println("❌ Failed to inject genesis accounts!")
		tmOs.Exit(err.Error())
	}

	authState.Accounts = accounts

	var rawAuthState bytes.Buffer
	_ = marshaler.Marshal(&rawAuthState, authState)

	return rawAuthState.Bytes()
}

func GenerateAuthzState() []byte {
	authzState := authzTypes.DefaultGenesisState()

	var rawAuthzState bytes.Buffer
	_ = marshaler.Marshal(&rawAuthzState, authzState)

	return rawAuthzState.Bytes()
}

func GenerateBankState(chainID string, denom string) []byte {
	bankState := bankTypes.DefaultGenesisState()

	balances, err := InjectGenesisBalances(chainID, denom)
	if err != nil {
		fmt.Println("❌ Failed to inject genesis balances!")
		tmOs.Exit(err.Error())
	}

	bankState.Balances = balances

	bankState.Supply = GenerateBankSupply(denom)
	bankState.DenomMetadata = GenerateBankMetadata(denom)

	var rawBankState bytes.Buffer
	_ = marshaler.Marshal(&rawBankState, bankState)

	return rawBankState.Bytes()
}

func GenerateCapabilityState() []byte {
	capabilityState := capabilityTypes.DefaultGenesis()

	var rawCapabilityState bytes.Buffer
	_ = marshaler.Marshal(&rawCapabilityState, capabilityState)

	return rawCapabilityState.Bytes()
}

func GenerateCrisisState(denom string) []byte {
	crisisState := crisisTypes.DefaultGenesisState()

	// Set the crisis fee to 10,000 $KYVE.
	crisisState.ConstantFee = sdk.NewCoin(denom, sdk.NewIntFromUint64(10_000*1_000_000))

	var rawCrisisState bytes.Buffer
	_ = marshaler.Marshal(&rawCrisisState, crisisState)

	return rawCrisisState.Bytes()
}

func GenerateDistributionState() []byte {
	distributionState := distributionTypes.DefaultGenesisState()

	var rawDistributionState bytes.Buffer
	_ = marshaler.Marshal(&rawDistributionState, distributionState)

	return rawDistributionState.Bytes()
}

func GenerateEvidenceState() []byte {
	evidenceState := evidenceTypes.DefaultGenesisState()

	var rawEvidenceState bytes.Buffer
	_ = marshaler.Marshal(&rawEvidenceState, evidenceState)

	return rawEvidenceState.Bytes()
}

func GenerateFeeGrantState() []byte {
	feeGrantState := feeGrantTypes.DefaultGenesisState()

	var rawFeeGrantState bytes.Buffer
	_ = marshaler.Marshal(&rawFeeGrantState, feeGrantState)

	return rawFeeGrantState.Bytes()
}

func GenerateGenUtilState(chainID string, unsafe bool) []byte {
	genUtilState, err := InjectGenesisTransactions(chainID, unsafe)
	if err == nil {
		fmt.Println(fmt.Sprintf("📝 Injected %d genesis transaction(s) ...", len(genUtilState.GenTxs)))
	} else {
		fmt.Println("❌ Failed to inject genesis transactions!")
		tmOs.Exit(err.Error())
	}

	var rawGenUtilState bytes.Buffer
	_ = marshaler.Marshal(&rawGenUtilState, genUtilState)

	return rawGenUtilState.Bytes()
}

func GenerateGovState(denom string) []byte {
	govState := govTypes.DefaultGenesisState()

	// Set the minimum deposit to 25,000 $KYVE.
	govState.DepositParams.MinDeposit = sdk.NewCoins(
		sdk.NewCoin(denom, sdk.NewIntFromUint64(25_000*1_000_000)),
	)
	// Set the deposit period to 5 minutes.
	fiveMinutes := 5 * time.Minute
	govState.DepositParams.MaxDepositPeriod = &fiveMinutes
	// Set the voting period to 1 hour.
	oneHour := time.Hour
	govState.VotingParams.VotingPeriod = &oneHour

	var rawGovState bytes.Buffer
	_ = marshaler.Marshal(&rawGovState, govState)

	return rawGovState.Bytes()
}

func GenerateGroupState() []byte {
	groupState := groupTypes.NewGenesisState()

	var rawGroupState bytes.Buffer
	_ = marshaler.Marshal(&rawGroupState, groupState)

	return rawGroupState.Bytes()
}

func GenerateMintState(denom string) []byte {
	goalBonded := sdk.MustNewDecFromStr("0.292")
	// NOTE: This is assuming 6-second block times.
	blocksPerYear := uint64(365.25 * 24 * 60 * 60 / 6)

	minter := mintTypes.InitialMinter(sdk.ZeroDec())
	params := mintTypes.NewParams(
		denom, sdk.OneDec(), sdk.ZeroDec(), sdk.ZeroDec(), goalBonded, blocksPerYear,
	)
	mintState := mintTypes.NewGenesisState(minter, params)

	var rawMintState bytes.Buffer
	_ = marshaler.Marshal(&rawMintState, mintState)

	return rawMintState.Bytes()
}

func GenerateSlashingState() []byte {
	slashingState := slashingTypes.DefaultGenesisState()

	// NOTE: This is assuming 6-second block times.
	slashingState.Params.SignedBlocksWindow = int64(24 * 60 * 60 / 6)
	slashingState.Params.DowntimeJailDuration = 2 * time.Hour
	slashingState.Params.SlashFractionDowntime = sdk.MustNewDecFromStr("0.001")

	var rawSlashingState bytes.Buffer
	_ = marshaler.Marshal(&rawSlashingState, slashingState)

	return rawSlashingState.Bytes()
}

func GenerateStakingState(denom string) []byte {
	stakingState := stakingTypes.DefaultGenesisState()

	stakingState.Params.BondDenom = denom
	stakingState.Params.MinCommissionRate = sdk.MustNewDecFromStr("0.05")

	var rawStakingState bytes.Buffer
	_ = marshaler.Marshal(&rawStakingState, stakingState)

	return rawStakingState.Bytes()
}
