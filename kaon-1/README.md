# Kaon Testnet <sup>v1</sup>

> Genesis validator submissions were closed on Feb 3rd, 2023.

## Running a Validator

> IMPORTANT: This guide assumes you've already configured your validator
> instance when submitting a genesis transaction. A complete guide is coming
> later, post network launch.

### Step 1 — Upgrade `kyved`.

To correctly run your validator, ensure you're running the latest release of
the `kyved` binary. Note that you used an outdated binary to submit your
genesis transactions.

`darwin/amd64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_amd64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_amd64.tar.gz)

`darwin/arm64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_arm64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_arm64.tar.gz)

`linux/amd64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_amd64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_amd64.tar.gz)

`linux/arm64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_arm64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_arm64.tar.gz)

### Step 2 — Install `cosmovisor`.

**NOTE** — This assumes you have [Go](https://go.dev/) on your instance.

<!-- go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest -->
```shell
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest
```

### Step 3 — Initialise `cosmovisor` for Kaon.

Before initialising, we first need to export some required environment variables:

```shell
export DAEMON_NAME=kyved
export DAEMON_HOME=~/.kaon-1
```

Now, we can initialise `cosmovisor` using the following command:

```shell
cosmovisor init <path to kyved>
```

### Step 4 — Download & verify the Kaon genesis file.

```shell
curl https://raw.githubusercontent.com/KYVENetwork/networks/main/kaon-1/genesis.json > ~/.kaon-1/config/genesis.json
```

**NOTE** — This assumes you have [`sha256sum`](https://linux.die.net/man/1/sha256sum/) on your instance.

```shell
cd ~/.kaon-1/config
echo "3532166eb1605057f633ff577b4fc3e57a6dddc46498c5bc6f2f4e8ab0c756b8  genesis.json" | sha256sum -c
```

### Step 5 — Start `cosmovisor`.

```shell
cosmovisor start
```

<!--

## Becoming a Genesis Validator

### Step 1 — Install `kyved`.

For now, we are only providing pre-compiled `kyved` binaries. Note that we
might ship a new version of the binary before the network launch.

`darwin/amd64`: [https://files.kyve.network/chain/v0.8.0/kyved_darwin_amd64.tar.gz](
https://files.kyve.network/chain/v0.8.0/kyved_darwin_amd64.tar.gz)

`darwin/arm64`: [https://files.kyve.network/chain/v0.8.0/kyved_darwin_arm64.tar.gz](
https://files.kyve.network/chain/v0.8.0/kyved_darwin_arm64.tar.gz)

`linux/amd64`: [https://files.kyve.network/chain/v0.8.0/kyved_linux_amd64.tar.gz](
https://files.kyve.network/chain/v0.8.0/kyved_linux_amd64.tar.gz)

`linux/arm64`: [https://files.kyve.network/chain/v0.8.0/kyved_linux_arm64.tar.gz](
https://files.kyve.network/chain/v0.8.0/kyved_linux_arm64.tar.gz)

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

### Step 5 — Generate a genesis transaction.

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

### Step 6 — Submit your genesis transaction.

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

-->
