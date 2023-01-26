# Kaon Testnet <sup>v1</sup>

## Becoming a Genesis Validator

### Step 1 — Install `kyved`.

### Step 2 — Initialise `kyved` for Kaon.

```shell
kyved init <moniker> --home ~/.kaon-1
```

### Step 3 — Create or import a key.

```shell
kyved keys add <name> --home ~/.kaon-1
```

The above command will generate a new key for you to use. However, you can
easily import a previously generated mnemonic using the `--recover` flag. The
`kyved` binary also comes with Ledger support, which you can access with the
`--ledger` flag.

### Step 4 — Register your account.

In this step, we will need to register your account in the genesis file; that
way, you can generate a genesis transaction. Each genesis validator will be
allocated 1 $KYVE (`1_000_000 tkyve`) for initial staking.

```shell
kyved add-genesis-account <address> 1000000tkyve --home ~/.kaon-1
```

Please note that you can find the address above using the following command:

```shell
kyved keys show <name> --address --home ~/.kaon-1
```

### Step 4 — Generate a genesis transaction.

The following command will create and sign a genesis transaction, successfully
creating your validator on network launch. Again, please note that as your
account only has 1 $KYVE (`1_000_000 tkyve`) registered, you won't be able to
customise the initial stake. You can also specify additional parameters for
your validator, but we have included the required ones.

```shell
kyved gentx <name> 1000000tkyve \
  --chain-id kaon-1 \
  --home ~/.kaon-1 \
  --moniker <moniker> \
  --details "My validator description."
```

### Step 5 — Submit your genesis transaction.

You will want to create a fork of this repository
([`KYVENetwork/networks`](https://github.com/KYVENetwork/networks/fork)) to
submit your genesis transaction. The above step should've given you a signed
genesis transaction that you will want to put into the
`./kaon-1/gentxs/<moniker>.json` file. Once you have completed this, please
open a PR, and the KYVE core team will review your submission as soon as
possible.

Please note that your address and your validator address are required when
submitting your PR. This will help with the foundation delegation program. You
can obtain your validator address with the following command (it will be in the
`Bech32 Val` section):

```shell
kyved debug addr <address>
```
