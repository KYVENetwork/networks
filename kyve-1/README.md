# KYVE Mainnet <sup>v1</sup>

## Becoming a Genesis Validator

### Step 1 — Install `kyved`.

### RECOMMENDED: Building from source.

> NOTE: You are required to have Go 1.19.x installed on your instance.

First, we'll want to clone the source code and check out the `v1.0.0-rc0`
release tag:

```shell
git clone https://github.com/KYVENetwork/chain
cd chain
git fetch
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

Additionally, to make network interactions via daemon easier, we're going to
configure the Chain ID globally:

```shell
kyved config chain-id kyve-1
```
