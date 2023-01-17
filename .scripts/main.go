package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	tmTypes "github.com/tendermint/tendermint/types"
	// Bank
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type AppState struct {
	BankState json.RawMessage `json:"bank"`
}

func main() {
	marshaler := jsonpb.Marshaler{Indent: "  "}

	chainID := flag.String("chain-id", "", "")
	flag.Parse()

	// x/bank
	bankState := bankTypes.DefaultGenesisState()
	var rawBankState bytes.Buffer
	_ = marshaler.Marshal(&rawBankState, bankState)

	appState := AppState{BankState: rawBankState.Bytes()}
	rawAppState, _ := json.Marshal(appState)

	genesis := tmTypes.GenesisDoc{
		// GenesisTime:     tmTime.Now(),
		ChainID:         *chainID,
		InitialHeight:   1,
		ConsensusParams: tmTypes.DefaultConsensusParams(),
		Validators:      nil,
		AppState:        json.RawMessage(rawAppState),
	}

	_ = genesis.SaveAs(fmt.Sprintf("../%s/genesis.json", *chainID))
}
