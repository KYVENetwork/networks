package main

import (
	"bytes"

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

	var rawDelegationState bytes.Buffer
	_ = marshaler.Marshal(&rawDelegationState, delegationState)

	return rawDelegationState.Bytes()
}

func GenerateGlobalState() []byte {
	globalState := globalTypes.DefaultGenesisState()

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

	var rawStakersState bytes.Buffer
	_ = marshaler.Marshal(&rawStakersState, stakersState)

	return rawStakersState.Bytes()
}

func GenerateTeamState() []byte {
	teamState := teamTypes.DefaultGenesis()

	var rawTeamState bytes.Buffer
	_ = marshaler.Marshal(&rawTeamState, teamState)

	return rawTeamState.Bytes()
}
