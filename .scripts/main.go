package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	kyveApp "github.com/KYVENetwork/chain/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genUtilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/gogo/protobuf/jsonpb"
	tmOs "github.com/tendermint/tendermint/libs/os"
	tmTypes "github.com/tendermint/tendermint/types"
)

var marshaler = jsonpb.Marshaler{
	EmitDefaults: true, Indent: "  ", OrigName: true,
}

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
	InitSDKConfig("kyve")

	chainID := flag.String("chain-id", "kyve-1", "")
	denom := flag.String("denom", "ukyve", "")
	startTime := flag.Int64("start-time", 1678786860, "")
	flag.Parse()

	fmt.Println(fmt.Sprintf("🤖 Creating genesis for %s ...", *chainID))
	fmt.Println(fmt.Sprintf("💰 Using %s as the global denom ...", *denom))

	appState := AppState{
		// NOTE: x/params is left as null intentionally.
		// NOTE: x/upgrade & x/vesting have been assigned to {} per Tendermint standard.
		AuthState:         GenerateAuthState(),
		AuthzState:        GenerateAuthzState(),
		BankState:         GenerateBankState(),
		CapabilityState:   GenerateCapabilityState(),
		CrisisState:       GenerateCrisisState(*denom),
		DistributionState: GenerateDistributionState(),
		EvidenceState:     GenerateEvidenceState(),
		FeeGrantState:     GenerateFeeGrantState(),
		GenUtilState:      GenerateGenUtilState(*chainID),
		GovState:          GenerateGovState(*denom),
		GroupState:        GenerateGroupState(),
		MintState:         GenerateMintState(*denom),
		SlashingState:     GenerateSlashingState(),
		StakingState:      GenerateStakingState(*denom),
		UpgradeState:      []byte("{}"),
		VestingState:      []byte("{}"),

		IBCState:         GenerateIBCState(),
		IBCFeeState:      GenerateIBCFeeState(),
		IBCTransferState: GenerateIBCTransferState(),
		ICAState:         GenerateICAState(),

		// NOTE: x/query is left as null intentionally.
		BundlesState:    GenerateBundlesState(),
		DelegationState: GenerateDelegationState(),
		GlobalState:     GenerateGlobalState(),
		PoolState:       GeneratePoolState(),
		StakersState:    GenerateStakersState(),
		TeamState:       GenerateTeamState(),
	}
	rawAppState, _ := json.Marshal(appState)

	genesis := tmTypes.GenesisDoc{
		GenesisTime: time.Unix(*startTime, 0),
		ChainID:     *chainID,
		AppState:    json.RawMessage(rawAppState),
	}

	validateErr := genesis.ValidateAndComplete()
	if validateErr != nil {
		fmt.Println("❌ Failed to validate genesis!")
		tmOs.Exit(validateErr.Error())
	}

	saveErr := genesis.SaveAs(fmt.Sprintf("../%s/genesis.json", *chainID))
	if saveErr != nil {
		fmt.Println("❌ Failed to save genesis file!")
		tmOs.Exit(saveErr.Error())
	} else {
		fmt.Println("✅ Completed genesis creation!")
	}
}

func InjectGenesisTransactions(chainID string) (*genUtilTypes.GenesisState, error) {
	dir, dirErr := os.ReadDir(fmt.Sprintf("../%s/gentxs", chainID))
	if dirErr != nil {
		return nil, dirErr
	}

	txDecoder := kyveApp.MakeEncodingConfig().TxConfig.TxJSONDecoder()
	txEncoder := kyveApp.MakeEncodingConfig().TxConfig.TxJSONEncoder()

	var genTxs []sdk.Tx

	for _, entry := range dir {
		file, _ := os.ReadFile(fmt.Sprintf("../%s/gentxs/%s", chainID, entry.Name()))

		tx, err := genUtilTypes.ValidateAndGetGenTx(file, txDecoder)
		if err == nil {
			genTxs = append(genTxs, tx)
		}
	}

	return genUtilTypes.NewGenesisStateFromTx(txEncoder, genTxs), nil
}
