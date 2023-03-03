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

After downloading the binary, you can verify that the sha256 hashes match.

`kyved_darwin_amd64.tar.gz` : `adca1016babd27c5f23ad40caf379d884556ea647c0d54fa01918c17b26803fb`\
`kyved_darwin_arm64.tar.gz` : `4c3c297ed6cc924fdc37f3a409f85837c71443ba12f6ff33163d123279fa2f36`\
`kyved_linux_amd64.tar.gz` : `15aa68a33a3427c8769613e6e433a4cf1d84308e0417a607fe59d10a830587af`\
`kyved_linux_arm64.tar.gz` : `6c92a5de44be1b450e82f049e034bfe0676f1c707cede0395844ad0183890a4d`

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
cosmovisor run start --home ~/.kaon-1
```
