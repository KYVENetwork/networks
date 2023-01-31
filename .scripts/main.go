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
	tmProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmTypes "github.com/tendermint/tendermint/types"

	// Auth
	authSigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// Bank
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// GenUtil
	genUtilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	// Staking
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	// Vesting
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
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
	dateString := flag.String("start-time", "2023-03-14 09:41:00", "")
	flag.Parse()

	startTime, err := time.Parse("2006-01-02 15:04:05", *dateString)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("🤖 Creating genesis for %s ...", *chainID))
	fmt.Println(fmt.Sprintf("💰 Using %s as the global denom ...", *denom))

	appState := AppState{
		// NOTE: x/params is left as null intentionally.
		// NOTE: x/upgrade & x/vesting have been assigned to {} per Tendermint standard.
		AuthState:         GenerateAuthState(*chainID, *denom),
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
		GenesisTime:     startTime,
		ChainID:         *chainID,
		ConsensusParams: GenerateConsensusParams(),
		AppState:        json.RawMessage(rawAppState),
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

func GenerateConsensusParams() *tmProto.ConsensusParams {
	return tmTypes.DefaultConsensusParams()
}

func InjectGenesisAccounts(chainID string, denom string) ([]*codecTypes.Any, error) {
	rawFile, openErr := os.Open(fmt.Sprintf("../%s/accounts.csv", chainID))
	if openErr != nil {
		return nil, openErr
	}

	file, readErr := csv.NewReader(rawFile).ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	var accounts []*codecTypes.Any

	for _, row := range file[1:] {
		// [ADDRESS] [AMOUNT]
		// NOTE: All addresses that aren't parsable are skipped.
		address, err := sdk.AccAddressFromBech32(row[0])
		if err != nil {
			continue
		}
		baseAccount := authTypes.NewBaseAccountWithAddress(address)

		var account authTypes.AccountI = baseAccount
		if row[3] != "" && row[4] != "" {
			var vestingAmount math.Int
			if row[2] == "" {
				vestingAmount, _ = math.NewIntFromString(row[1])
			} else {
				vestingAmount, _ = math.NewIntFromString(row[2])
			}

			vestingCoins := sdk.NewCoins(sdk.NewCoin(denom, vestingAmount))

			vestingStart, _ := time.Parse("2006-01-02", row[3])
			vestingEnd, _ := time.Parse("2006-01-02", row[4])

			account = vestingTypes.NewContinuousVestingAccount(
				baseAccount, vestingCoins, vestingStart.Unix(), vestingEnd.Unix(),
			)
		}

		rawAccount, err := codecTypes.NewAnyWithValue(account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, rawAccount)
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

	for _, row := range file[1:] {
		// [ADDRESS] [AMOUNT]
		// NOTE: All addresses that aren't parsable are treated as module accounts.
		address, err := sdk.AccAddressFromBech32(row[0])
		if err != nil {
			address = authTypes.NewModuleAddress(row[0])
		}

		amount, _ := math.NewIntFromString(row[1])
		coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

		balance := bankTypes.Balance{Address: address.String(), Coins: coins}
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

		tx, err := ValidateAndGetGenTx(file, txDecoder)
		if err != nil {
			continue
		}

		if VerifySignature(chainID, tx) {
			genTxs = append(genTxs, tx)
		}
	}

	return genUtilTypes.NewGenesisStateFromTx(txEncoder, genTxs), nil
}

// ValidateAndGetGenTx is inspired by:
// https://github.com/cosmos/cosmos-sdk/blob/release/v0.46.x/x/genutil/types/genesis_state.go#L108
func ValidateAndGetGenTx(genTx json.RawMessage, txJSONDecoder sdk.TxDecoder) (authSigning.SigVerifiableTx, error) {
	rawTx, err := txJSONDecoder(genTx)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gentx: %s, error: %s", genTx, err)
	}

	tx, ok := rawTx.(authSigning.SigVerifiableTx)
	if !ok {
		return nil, fmt.Errorf("invalid GenTx")
	}

	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		if sdk.MsgTypeURL(msg) != sdk.MsgTypeURL(&bankTypes.MsgMultiSend{}) &&
			sdk.MsgTypeURL(msg) != sdk.MsgTypeURL(&stakingTypes.MsgCreateValidator{}) &&
			sdk.MsgTypeURL(msg) != sdk.MsgTypeURL(&stakingTypes.MsgDelegate{}) {
			return nil, fmt.Errorf("unexpected GenTx message type; expected: MsgSend, MsgCreateValidator, or MsgDelegate, got: %T", msg)
		}

		if err := msg.ValidateBasic(); err != nil {
			return nil, fmt.Errorf("invalid GenTx '%s': %s", msg, err)
		}
	}

	return tx, nil
}

// VerifySignature is inspired by:
// https://github.com/cosmos/cosmos-sdk/blob/release/v0.46.x/x/auth/signing/verify.go#L14
func VerifySignature(chainID string, tx authSigning.SigVerifiableTx) bool {
	signatures, err := tx.GetSignaturesV2()
	if err != nil || len(signatures) != 1 {
		return false
	}

	pubKey := signatures[0].PubKey
	signerData := authSigning.SignerData{
		Address:       pubKey.Address().String(),
		ChainID:       chainID,
		AccountNumber: 0,
		Sequence:      0,
		PubKey:        pubKey,
	}

	err = authSigning.VerifySignature(
		pubKey, signerData, signatures[0].Data, kyveApp.MakeEncodingConfig().TxConfig.SignModeHandler(), tx,
	)
	if err != nil {
		return false
	}

	return true
}
