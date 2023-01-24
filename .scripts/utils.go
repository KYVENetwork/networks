package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func InitSDKConfig(accountAddressPrefix string) {
	accountPubKeyPrefix := accountAddressPrefix + "pub"
	validatorAddressPrefix := accountAddressPrefix + "valoper"
	validatorPubKeyPrefix := accountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := accountAddressPrefix + "valcons"
	consNodePubKeyPrefix := accountAddressPrefix + "valconspub"

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(accountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}

// ========== CONSTANTS ==========

func GenerateBankMetadata(denom string) []bankTypes.Metadata {
	units := []*bankTypes.DenomUnit{
		{
			Denom:    denom,
			Exponent: 0,
			Aliases:  []string{"microkyve"},
		},
		{
			Denom:    "kyve",
			Exponent: 6,
		},
	}

	metadata := []bankTypes.Metadata{
		{
			Description: "The native utility token of the KYVE network.",
			DenomUnits:  units,
			Base:        denom,
			Display:     "kyve",
			Name:        "KYVE",
			Symbol:      "KYVE",
		},
	}

	return metadata
}

func GenerateBankSupply(denom string) sdk.Coins {
	return sdk.NewCoins(
		sdk.NewCoin(denom, sdk.NewIntFromUint64(1_000_000_000)),
	)
}
