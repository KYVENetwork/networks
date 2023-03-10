# KYVE Mainnet <sup>v1</sup>

> Genesis validator submissions were closed on Mar 8th, 2023.

## Running a Validator

> IMPORTANT: This guide assumes you've configured your validator
> instance when submitting a genesis transaction. A complete guide is coming
> later, post-network launch.

### Step 1 — Upgrade `kyved`.

**NOTE:** If you are building from source, you **DO NOT** need to specify the
`ENV` variable.

To correctly run your validator, ensure you're using the latest `kyved` binary
release. Note that you used an outdated binary to submit your genesis
transactions.

After downloading the prebuilt binaries from the [`v1.0.0`](https://github.com/KYVENetwork/chain/releases/tag/v1.0.0)
release page or building from source, you can verify that the sha256 hashes match.

`kyved_darwin_amd64.tar.gz` : `4f7ef862d11573ef98879a4463c0cc9cbb021fcedd510875b7123d4aecb56918`\
`kyved_darwin_arm64.tar.gz` : `cc2b2ee21fbda39961cc9f62ec3e2eac90b1ab42a911906255d674648df863a4`\
`kyved_linux_amd64.tar.gz` : `f594ead2340f055eec8801a4b38d3849bcd900f3ed820c8962458abfa45dd7ca`\
`kyved_linux_arm64.tar.gz` : `1f17d4deb5dcef24f4691e57ce621d946238bca455e7dce9f04a4648b3f771ab`

### Step 2 - Verify `kyved`.

You can verify that you are running the correct binary by running the following commands:

```shell
kyved version

# v1.0.0
```

```shell
kyved info

# Information about build variables:
# Version: v1.0.0
# Denom: ukyve
# Team-Foundation-Authority: kyve1xjpl57p7f49y5gueu7rlfytaw9ramcn5zhjy2g
# Team-BCP-Authority: kyve1fnh4kghr25tppskap50zk5j385pt65tyyjaraa
# Team-Allocation: 165,000,000.000000
# Team-TGE: 2023-03-14 15:03:14 +0100 CET
```

### Step 2 — Install `cosmovisor`.

**NOTE** — This assumes you have [Go](https://go.dev/) on your instance.

```shell
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest
```

### Step 3 — Initialise `cosmovisor` for KYVE.

Before initialising, we first need to export some required environment variables:

```shell
export DAEMON_NAME=kyved
export DAEMON_HOME=~/.kyve
```

Now, we can initialise `cosmovisor` using the following command:

```shell
cosmovisor init <path to kyved>
```

### Step 4 — Download & verify the KYVE genesis file.

```shell
curl https://raw.githubusercontent.com/KYVENetwork/networks/main/kyve-1/genesis.json > ~/.kyve/config/genesis.json
```

**NOTE** — This assumes you have [`sha256sum`](https://linux.die.net/man/1/sha256sum/) on your instance.

```shell
cd ~/.kyve/config
echo "1dc3ec916f49ef8c221851566aca12a3f914b23afb3ab35067fc8a8d5f59c2ee  genesis.json" | sha256sum -c
```

### Step 5 — Start `cosmovisor`.

```shell
cosmovisor run start
```

<!--

## Becoming a Genesis Validator

### Step 1 — Install `kyved`.

### RECOMMENDED: Building from source.

> NOTE: You are required to have Go 1.19.x installed on your instance.

First, we'll want to clone the source code and check out the `v1.0.0-rc0`
release tag:

```shell
git clone https://github.com/KYVENetwork/chain
cd chain
git fetch --tags
git checkout v1.0.0-rc0
```

Now that we have checked out the correct tag, we can go ahead and run the
following:

```shell
make install
```

If you've configured your Go paths correctly, you will now have `kyved` in your
global path. To double-check that you're running the correct version, run the
following:

```shell
kyved version
# 1.0.0-rc0
```

### Prebuilt binaries.

`darwin/amd64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_amd64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_amd64.tar.gz)

`darwin/arm64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_arm64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_darwin_arm64.tar.gz)

`linux/amd64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_amd64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_amd64.tar.gz)

`linux/arm64`: [https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_arm64.tar.gz](
https://files.kyve.network/chain/v1.0.0-rc0/kyved_linux_arm64.tar.gz)

### Step 2 — Initialise `kyved`.

```shell
kyved init <moniker>
```

Note this will initialise your validator in the default `~/.kyve` home 
directory.

Additionally, to make network interactions via the daemon easier, we're going
to configure the Chain ID globally:

```shell
kyved config chain-id kyve-1
```

### Step 3 — Create or import a key.

```shell
kyved keys add <name>
```

The above command will generate a new key for you to use. However, you can
easily import a previously generated mnemonic using the `--recover` flag. The
`kyved` daemon also comes with Ledger support, which you can access with the
`--ledger` flag.

### Step 4 — Register your account.

In this step, we will need to register your account in the genesis file; that
way, you can generate a genesis transaction. Each genesis validator will be
allocated 1 $KYVE (`1_000_000 ukyve`) for initial staking.

```shell
kyved add-genesis-account <address> 1000000ukyve
```

Please note that you can find the address above using the following command:

```shell
kyved keys show <name> --address
```

### Step 5 — Generate a genesis transaction.

The following command will create and sign a genesis transaction, successfully
creating your validator on network launch. Again, please note that as your
account only has 1 $KYVE (`1_000_000 ukyve`) registered, you won't be able to
customise the initial stake. You can also specify additional parameters for
your validator, but we have included the required ones.

```shell
kyved gentx <name> 1000000ukyve \
  --moniker <moniker> \
  --identity <identity> \
  --details <description> \
  --security-contact <security-contact> \
  --website <website>
```

### Step 6 — Submit your genesis transaction.

You will want to create a fork of this repository
([`KYVENetwork/networks`](https://github.com/KYVENetwork/networks/fork)) to
submit your genesis transaction. The above step should've given you a signed
genesis transaction that you will want to put into the
`./kyve-1/gentxs/<moniker>.json` file. Once you have completed this, please
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
