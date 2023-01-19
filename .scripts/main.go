package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"cosmossdk.io/math"
	kyveApp "github.com/KYVENetwork/chain/app"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
	tmOs "github.com/tendermint/tendermint/libs/os"
	tmTypes "github.com/tendermint/tendermint/types"

	// Bank
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// GenUtil
	genUtilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
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
		AuthState:         GenerateAuthState(*chainID),
		AuthzState:        GenerateAuthzState(),
		BankState:         GenerateBankState(*chainID, *denom),
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

func InjectGenesisAccounts(chainID string) ([]*codecTypes.Any, error) {
	rawFile, openErr := os.Open(fmt.Sprintf("../%s/accounts.csv", chainID))
	if openErr != nil {
		return nil, openErr
	}

	file, readErr := csv.NewReader(rawFile).ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	var accounts []*codecTypes.Any

	for range file {
		// [ADDRESS] [AMOUNT]
	}

	return accounts, nil
}

func InjectGenesisBalances(chainID string, denom string) ([]bankTypes.Balance, error) {
	rawFile, openErr := os.Open(fmt.Sprintf("../%s/accounts.csv", chainID))
	if openErr != nil {
		return nil, openErr
	}

	file, readErr := csv.NewReader(rawFile).ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	var balances []bankTypes.Balance

	for _, row := range file {
		// [ADDRESS] [AMOUNT]
		amount, _ := math.NewIntFromString(row[1])
		coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

		balance := bankTypes.Balance{Address: row[0], Coins: coins}
		balances = append(balances, balance)
	}

	return balances, nil
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
