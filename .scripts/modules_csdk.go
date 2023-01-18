package main

import (
	"bytes"

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
	// GenUtil
	genUtilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
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

func GenerateAuthState() []byte {
	authState := authTypes.DefaultGenesisState()

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

func GenerateBankState() []byte {
	bankState := bankTypes.DefaultGenesisState()

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

func GenerateCrisisState(_ string) []byte {
	crisisState := crisisTypes.DefaultGenesisState()

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

func GenerateGenUtilState() []byte {
	genUtilState := genUtilTypes.DefaultGenesisState()

	var rawGenUtilState bytes.Buffer
	_ = marshaler.Marshal(&rawGenUtilState, genUtilState)

	return rawGenUtilState.Bytes()
}

func GenerateGovState(_ string) []byte {
	govState := govTypes.DefaultGenesisState()

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
	mintState := mintTypes.DefaultGenesisState()

	mintState.Params.MintDenom = denom

	var rawMintState bytes.Buffer
	_ = marshaler.Marshal(&rawMintState, mintState)

	return rawMintState.Bytes()
}

func GenerateSlashingState() []byte {
	slashingState := slashingTypes.DefaultGenesisState()

	var rawSlashingState bytes.Buffer
	_ = marshaler.Marshal(&rawSlashingState, slashingState)

	return rawSlashingState.Bytes()
}

func GenerateStakingState(denom string) []byte {
	stakingState := stakingTypes.DefaultGenesisState()

	stakingState.Params.BondDenom = denom

	var rawStakingState bytes.Buffer
	_ = marshaler.Marshal(&rawStakingState, stakingState)

	return rawStakingState.Bytes()
}
