package main

import (
	"bytes"

	// IBC
	ibcTypes "github.com/cosmos/ibc-go/v5/modules/core/types"
	// IBC Fee
	ibcFeeTypes "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/types"
	// IBC Transfer
	ibcTransferTypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	// ICA
	icaTypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
)

func GenerateIBCState() []byte {
	ibcState := ibcTypes.DefaultGenesisState()

	var rawIBCState bytes.Buffer
	_ = marshaler.Marshal(&rawIBCState, ibcState)

	return rawIBCState.Bytes()
}

func GenerateIBCFeeState() []byte {
	ibcFeeState := ibcFeeTypes.DefaultGenesisState()

	var rawIBCFeeState bytes.Buffer
	_ = marshaler.Marshal(&rawIBCFeeState, ibcFeeState)

	return rawIBCFeeState.Bytes()
}

func GenerateIBCTransferState() []byte {
	ibcTransferState := ibcTransferTypes.DefaultGenesisState()

	ibcTransferState.Params.SendEnabled = false
	ibcTransferState.Params.ReceiveEnabled = false

	var rawIBCTransferState bytes.Buffer
	_ = marshaler.Marshal(&rawIBCTransferState, ibcTransferState)

	return rawIBCTransferState.Bytes()
}

func GenerateICAState() []byte {
	icaState := icaTypes.DefaultGenesis()

	var rawICAState bytes.Buffer
	_ = marshaler.Marshal(&rawICAState, icaState)

	return rawICAState.Bytes()
}
