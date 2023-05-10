#!/bin/bash
BASEDIR=$(dirname "$0")
. "$BASEDIR"/../shared.sh

#Register Sequencer
DESCRIPTION="{\"Moniker\":\"$MONIKER\",\"Identity\":\"\",\"Website\":\"\",\"SecurityContact\":\"\",\"Details\":\"\"}";
SEQ_PUB_KEY="$($EXECUTABLE furyint show-sequencer --home $ROLLAPP_CHAIN_DIR)"

$SETTLEMENT_EXECUTABLE tx sequencer create-sequencer "$SEQ_PUB_KEY" "$ROLLAPP_ID" "$DESCRIPTION" \
  --from "$KEY_NAME_FURY" \
  --chain-id "$SETTLEMENT_CHAIN_ID" \
  --keyring-backend test \
  --keyring-dir "$KEYRING_PATH" \
  --broadcast-mode block \
  --node "$SETTLEMENT_RPC"