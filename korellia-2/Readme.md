# Korellia-2

Korellia is the official devnet. Its purpose is to help protocol developers
to get started as easily as possible. 

## Setup

- The network is run by validators from the KYVE Core Team which will always 
  keep the majority of the stake. However, there will still be 100 slots in 
  total so that it's possible for others to join.
- The frontend for Korellia is hosted at: https://app.korellia.kyve.network/#/
- A block explorer can be found here: https://explorer.kyve.network/korellia
- A pool can be created through the dedicated pools page. The backend will take
  care of putting the correct proposals on chain and voting.
- The pool creators can always update their own pools.
- $KYVE for Korellia can be claimed via the faucet: https://app.korellia.kyve.network/#/faucet

## Regenesis

Korellia has a long history of state breaking changes. It was the first KYVE
(test) chain launched and was used for the Incentivised Testnet in 2022.
Mainnet launch was successful and a working testnet (Kaon) is in place,
It was now time for a Korellia regenesis to clean up the state. Especially the 
relaunch is needed to switch from 9 decimals to 6 decimals. 
The following state will be kept / migrated:

- All balances of users
- All finalised bundles
- All protocol stakers
- All pools

## Preparation

First, obtain the previous Korellia genesis file and verify its hash.

```shell
wget -qO- https://arweave.net/-hZr0b1v0Qwkm88A1g1wTOiptqeA0LQi2KGBP4VPCBg | gzip -d > korellia.json
echo "b89fa04424cc3637b529c25d2db53fa95b971481f6285bb4bf10efc02b7a520a korellia.json" | sha256sum -c
```

Then run the python-script for generating the new genesis out of the existing 
genesis file.

```shell
python3 KorelliaRegenesis.py korellia.json genesis.json
```

## Joining the network

A prebuilt genesis file can be obtained via 

```shell
wget https://files.kyve.network/korellia-2/genesis.json
echo "c9c05363d5c535a1b2a9ff51ec63c878fad26e081f1fecf1d011a92dbbeeabbf genesis.json" | sha256sum -c
```

Obtain the binary for linux via:
```shell
wget https://files.kyve.network/korellia-2/1.4.0/kyved_kaon_linux_amd64.tar.gz
echo "6709369bbb6db4e3bc1b6445f671fbf2f0279668cb5a5f2ddef5b4674bfd099e kyved_kaon_linux_amd64.tar.gz | sha256sum -c
```

**Other prebuilt binaries**

| Arch         | URL                                                                        | SHA256                                                           |
|--------------|----------------------------------------------------------------------------|------------------------------------------------------------------|
| linux/arm64  | https://files.kyve.network/korellia-2/1.4.0/kyved_kaon_linux_arm64.tar.gz  | 3f6e20177fb48419b1b174aeed79d3129ed16aaca94438dd048a20b586f98252 |
| darwin/arm64 | https://files.kyve.network/korellia-2/1.4.0/kyved_kaon_darwin_arm64.tar.gz | 65a3764468ed76fa95f419848f95da54af9a1dc129eed5a178395a85186c5374 |
| darwin/amd64 | https://files.kyve.network/korellia-2/1.4.0/kyved_kaon_darwin_amd64.tar.gz | 00be5e6c0cd8247101d935de4adcd706d1f310411f86f58addd8a085c814d48b |


Use the following peers to connect to the network:
```text
2e6a696aac4a5b335d3bdc2fc90f2b602b7cbf11@35.80.149.217:26656
3de02301e9e899ffa4d61216f3fb3bca6f2d355f@52.58.250.62:26656
```

For the initial startup make sure to have at least 24G of RAM. If you don't
have enough RAM use a swapfile.
