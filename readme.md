<h1 align="center">Gridiron Rollapp</h1>

![banner](https://user-images.githubusercontent.com/109034310/204804891-bdc0f7bc-4b17-4b4a-99ff-25153d3887ee.jpg)
[![license](https://img.shields.io/github/license/cosmos/cosmos-sdk.svg#thumbnail)](https://github.com/gridironxyz/rdk/blob/main/LICENSE)

Gridiron RDK, which stands for *RollApp Development Kit* is based on the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) framework, modified and enhanced for building autonomous RollApps (app-specific-rollups) on top of the [Gridiron Hub](https://github.com/gridironxyz/gridiron).

The RDK provides the following capabilities ***on top*** of the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) framework:

* The RDK is coupled with the [Furyint](https://github.com/gridironxyz/furyint) client to form RollApp's blazing speed consensus and networking layer, while ***the Gridiron Hub is securing the rollapp***
* Custom modules that convert a cosmos based PoS (proof-of-stake) chain to a rollapp
* wasm and EVM support (based on CosmWasm and Ethermint)

### Learn more

For more information about Gridiron RollApps please visit the [documentation center](https://docs.gridiron.xyz/)

To learn how the Cosmos SDK works from a high-level perspective, see the Cosmos SDK [High-Level Intro](https://docs.cosmos.network/main/intro/overview.html).

If you want to get started quickly and learn how to build on top of Cosmos SDK, visit [Cosmos SDK Tutorials](https://tutorials.cosmos.network). You can also fork the tutorial's repository to get started building your own Cosmos SDK application.

---

This repository hosts `rollappd`, the first implementation of a gridiron rollapp.

**Note**: Requires [Go 1.18](https://go.dev/)

## Quick guide

Get started with [building RollApps](https://docs.gridiron.xyz/develop/get-started/setup)

## Installing / Getting started

Build and install the ```rollappd``` binary:

```shell
make install
```

### Initial configuration

Set custom configuration params at `scripts/shared.sh`
This will initialize the rollapp:

```shell
sh scripts/init_rollapp.sh
```

### Run rollapp

```shell
sh scripts/run_rollapp.sh
```

## Run a rollapp with local settlement node

### Run local gridiron hub node

Follow the instructions on [Gridiron Hub docs](https://docs.gridiron.xyz/develop/get-started/run-base-layers) to run local gridiron hub node

### Configure the rollapp

Set the settlement configuration params in `scripts/shared.sh` or by export

```shell
export SETTLEMENT_LAYER="gridiron"
export SETTLEMENT_RPC="tcp://127.0.0.1:36657"

sh scripts/init_rollapp.sh
```

### Register rollapp on settlement

```shell
sh scripts/settlement/register_rollapp_to_hub.sh
```

### Register sequencer for rollapp on settlement

```shell
sh scripts/settlement/register_sequencer_to_hub.sh
```

### Run rollapp

```shell
sh scripts/run_rollapp.sh
```

## Running EVM-based rollapp

:construction:  To run an EVM-based rollapp, one should:

1. Build and EVM binary instead of default binary:

    ```shell
    make install_evm
    ```

    This will build and install the ```rollapp_evm``` binary

2. EVM-based configuration:

    Uncomment the **EVM section** in `scripts/shared.sh`, **then** initializing the rollapp

    ```shell
    sh scripts/init_rollapp.sh
    ```

## Establish IBC channel between hub and rollapp

The following script will create all the dependencies for IBC channel between the hub and the rollapp.
It will create dedicated accounts for the relayer on both the hub and the rollapp, and transfer some funds to them from the genesis accounts.

```shell
sh scripts/ibc/setup_ibc.sh
```

after it finishes (it might take few mins), run the relayer:

```shell
sh scripts/run_relayer.sh
```

To run ibc-transfers between rollapp and the hub,
first check and set the connection name:

```shell
//check the connectionID of the rollapp and the hub on the active path
rly paths show hub-rollapp --json | jq '.chains.src'
rly paths show hub-rollapp --json | jq '.chains.dst'

//check the channel_ID based on the connectionID
rollappd q ibc channel connections <connectionID from rly command> -o json | jq '.channels[0].channel_id'
furyd q ibc channel connections <connectionID from rly command> -o json | jq '.channels[0].channel_id'

//Use the above result
export ROLLAPP_CHANNEL_NAME=<channel_id>
export HUB_CHANNEL_NAME=<channel_id>
```

Now you can do ibc transfers

```shell
sh scripts/ibc/ibc_transfer.sh [arg]

Available:
-q:         query balances of local-user on hub and rol-user on rollapp
rol2hub:    ibc-transfer of 5555urap to local-user from rol-user
hub_back:   transfer back the tokens from the hub to the rollapp
hub2rol:    ibc-transfer of 5555fury to rol-user from local-user
hub_back:   transfer back the tokens from the hub to the rollapp
```

### Run rollapp with IBC native token

To run a rollapp based on foreign IBC token, set the following when initializing the rollapp:

```shell
DENOM=IBC/<denom trace>
TOKEN_AMOUNT = 0$DENOM
```

## Running multiple rollapp instances locally

Run the first rollapp as described above.

For the 2nd rollapp, run the following in a new tab:

```shell
export CHAIN_ID=rollapp2
export CHAIN_DIR="$HOME/.rollapp2"
export ROLLAPP_ID=rollapp2
export RPC_PORT="0.0.0.0:27667"
export P2P_PORT="0.0.0.0:27668"
export GRPC_PORT="0.0.0.0:9180"
export GRPC_WEB_PORT="0.0.0.0:9181"
```

Then run the scripts as described in the readme

## Developers guide

TODO
