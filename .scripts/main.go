package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	tmTypes "github.com/tendermint/tendermint/types"

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

type AppState struct {
	AuthState         json.RawMessage `json:"auth"`
	AuthzState        json.RawMessage `json:"authz"`
	BankState         json.RawMessage `json:"bank"`
	BundlesState      json.RawMessage `json:"bundles"`
	CapabilityState   json.RawMessage `json:"capability"`
	CrisisState       json.RawMessage `json:"crisis"`
	DelegationState   json.RawMessage `json:"delegation"`
	DistributionState json.RawMessage `json:"distribution"`
	EvidenceState     json.RawMessage `json:"evidence"`
	FeeGrantState     json.RawMessage `json:"feegrant"`
	IBCFeeState       json.RawMessage `json:"feeibc"`
	GenUtilState      json.RawMessage `json:"genutil"`
	GlobalState       json.RawMessage `json:"global"`
	GovState          json.RawMessage `json:"gov"`
	GroupState        json.RawMessage `json:"group"`
	IBCState          json.RawMessage `json:"ibc"`
	ICAState          json.RawMessage `json:"interchainaccounts"`
	MintState         json.RawMessage `json:"mint"`
	ParamsState       json.RawMessage `json:"params"`
	PoolState         json.RawMessage `json:"pool"`
	QueryState        json.RawMessage `json:"query"`
	SlashingState     json.RawMessage `json:"slashing"`
	StakersState      json.RawMessage `json:"stakers"`
	StakingState      json.RawMessage `json:"staking"`
	TeamState         json.RawMessage `json:"team"`
	IBCTransferState  json.RawMessage `json:"transfer"`
	UpgradeState      json.RawMessage `json:"upgrade"`
	VestingState      json.RawMessage `json:"vesting"`
}

func main() {
	chainID := flag.String("chain-id", "kyve-1", "")
	denom := flag.String("denom", "ukyve", "")
	flag.Parse()

	fmt.Println(fmt.Sprintf("🤖 Creating genesis for %s ...", *chainID))
	fmt.Println(fmt.Sprintf("💰 Using %s as the global denom ...", *denom))

	appState := AppState{
		AuthState:         generateAuthState(),
		AuthzState:        generateAuthzState(),
		BankState:         generateBankState(),
		CapabilityState:   generateCapabilityState(),
		CrisisState:       generateCrisisState(*denom),
		DistributionState: generateDistributionState(),
		EvidenceState:     generateEvidenceState(),
		FeeGrantState:     generateFeeGrantState(),
		GenUtilState:      generateGenUtilState(),
		GovState:          generateGovState(*denom),
		GroupState:        generateGroupState(),
		MintState:         generateMintState(*denom),
		// NOTE: x/params is empty on purpose.
		SlashingState: generateSlashingState(),
		StakingState:  generateStakingState(*denom),
		// TODO(@john): Look into x/upgrade state.
		// TODO(@john): Look into x/vesting state.

		IBCState:         GenerateIBCState(),
		IBCFeeState:      GenerateIBCFeeState(),
		IBCTransferState: GenerateIBCTransferState(),
		ICAState:         GenerateICAState(),

		BundlesState:    GenerateBundlesState(),
		DelegationState: GenerateDelegationState(),
		GlobalState:     GenerateGlobalState(),
		PoolState:       GeneratePoolState(),
		// NOTE: x/query is empty on purpose.
		StakersState: GenerateStakersState(),
		TeamState:    GenerateTeamState(),
	}
	rawAppState, _ := json.Marshal(appState)

	genesis := tmTypes.GenesisDoc{
		// GenesisTime:     tmTime.Now(),
		ChainID:         *chainID,
		InitialHeight:   1,
		ConsensusParams: tmTypes.DefaultConsensusParams(),
		Validators:      nil,
		AppState:        json.RawMessage(rawAppState),
	}

	// TODO(@john): Start using the `ValidateAndComplete` function provided.

	// TODO(@john): Catch error when saving.
	_ = genesis.SaveAs(fmt.Sprintf("../%s/genesis.json", *chainID))
	fmt.Println("✅ Completed genesis creation!")
}

// ========== Module Functions ==========

var marshaler = jsonpb.Marshaler{
	EmitDefaults: true, Indent: "  ", OrigName: true,
}

// x/auth
func generateAuthState() []byte {
	authState := authTypes.DefaultGenesisState()

	var rawAuthState bytes.Buffer
	_ = marshaler.Marshal(&rawAuthState, authState)

	return rawAuthState.Bytes()
}

// x/authz
func generateAuthzState() []byte {
	authzState := authzTypes.DefaultGenesisState()

	var rawAuthzState bytes.Buffer
	_ = marshaler.Marshal(&rawAuthzState, authzState)

	return rawAuthzState.Bytes()
}

// x/bank
func generateBankState() []byte {
	bankState := bankTypes.DefaultGenesisState()

	var rawBankState bytes.Buffer
	_ = marshaler.Marshal(&rawBankState, bankState)

	return rawBankState.Bytes()
}

// x/capability
func generateCapabilityState() []byte {
	capabilityState := capabilityTypes.DefaultGenesis()

	var rawCapabilityState bytes.Buffer
	_ = marshaler.Marshal(&rawCapabilityState, capabilityState)

	return rawCapabilityState.Bytes()
}

// x/crisis
func generateCrisisState(_ string) []byte {
	crisisState := crisisTypes.DefaultGenesisState()

	var rawCrisisState bytes.Buffer
	_ = marshaler.Marshal(&rawCrisisState, crisisState)

	return rawCrisisState.Bytes()
}

// x/distribution
func generateDistributionState() []byte {
	distributionState := distributionTypes.DefaultGenesisState()

	var rawDistributionState bytes.Buffer
	_ = marshaler.Marshal(&rawDistributionState, distributionState)

	return rawDistributionState.Bytes()
}

// x/evidence
func generateEvidenceState() []byte {
	evidenceState := evidenceTypes.DefaultGenesisState()

	var rawEvidenceState bytes.Buffer
	_ = marshaler.Marshal(&rawEvidenceState, evidenceState)

	return rawEvidenceState.Bytes()
}

// x/feegrant
func generateFeeGrantState() []byte {
	feeGrantState := feeGrantTypes.DefaultGenesisState()

	var rawFeeGrantState bytes.Buffer
	_ = marshaler.Marshal(&rawFeeGrantState, feeGrantState)

	return rawFeeGrantState.Bytes()
}

// x/genutil
func generateGenUtilState() []byte {
	genUtilState := genUtilTypes.DefaultGenesisState()

	var rawGenUtilState bytes.Buffer
	_ = marshaler.Marshal(&rawGenUtilState, genUtilState)

	return rawGenUtilState.Bytes()
}

// x/gov
func generateGovState(_ string) []byte {
	govState := govTypes.DefaultGenesisState()

	var rawGovState bytes.Buffer
	_ = marshaler.Marshal(&rawGovState, govState)

	return rawGovState.Bytes()
}

// x/group
func generateGroupState() []byte {
	groupState := groupTypes.NewGenesisState()

	var rawGroupState bytes.Buffer
	_ = marshaler.Marshal(&rawGroupState, groupState)

	return rawGroupState.Bytes()
}

// x/mint
func generateMintState(denom string) []byte {
	mintState := mintTypes.DefaultGenesisState()

	mintState.Params.MintDenom = denom

	var rawMintState bytes.Buffer
	_ = marshaler.Marshal(&rawMintState, mintState)

	return rawMintState.Bytes()
}

// x/slashing
func generateSlashingState() []byte {
	slashingState := slashingTypes.DefaultGenesisState()

	var rawSlashingState bytes.Buffer
	_ = marshaler.Marshal(&rawSlashingState, slashingState)

	return rawSlashingState.Bytes()
}

// x/staking
func generateStakingState(denom string) []byte {
	stakingState := stakingTypes.DefaultGenesisState()

	stakingState.Params.BondDenom = denom

	var rawStakingState bytes.Buffer
	_ = marshaler.Marshal(&rawStakingState, stakingState)

	return rawStakingState.Bytes()
}
