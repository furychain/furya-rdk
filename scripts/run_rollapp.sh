#!/bin/bash
BASEDIR=$(dirname "$0")
. "$BASEDIR"/shared.sh

# If the settlement layer is set to gridiron, use $SETTLEMENT_CONFIG otherwise use $SETTLEMENT_CONFIG_MOCK
if [ ! "$SETTLEMENT_LAYER" = "gridiron" ]; then
  echo "using mock settlement layer"
  SETTLEMENT_CONFIG="$SETTLEMENT_CONFIG_MOCK"
fi

if [ "$DA_LAYER" = "mock" ]; then
  echo "using mock DA layer"
  DA_LAYER_CONFIG="30s"
fi

# If aggregator is set to true pass the aggregator flag
if [ "$AGGREGATOR" = "true" ]; then
  AGGREGATOR_FLAG="--furyint.aggregator"
else
  AGGREGATOR_FLAG=""
fi

if [ "$EVM_ENABLED" ]; then
  EVM_FLAGS='--json-rpc.enable --json-rpc.api eth,txpool,personal,net,debug,web3,miner'
fi


$EXECUTABLE start $AGGREGATOR_FLAG \
  $EVM_FLAGS \
  --furyint.da_layer "$DA_LAYER" \
  --furyint.da_config "$DA_LAYER_CONFIG" \
  --furyint.settlement_layer "$SETTLEMENT_LAYER" \
  --furyint.settlement_config "$SETTLEMENT_CONFIG" \
  --furyint.block_batch_size "$BATCH_SIZE" \
  --furyint.namespace_id "$NAMESPACE_ID" \
  --furyint.block_time "$BLOCK_TIME" \
  --home "$ROLLAPP_CHAIN_DIR" \
  --log_level "$LOG_LEVEL" \
  --log-file "$LOG_FILE_PATH" \
  --max-log-size "$MAX_LOG_SIZE" \
  --module-log-level-override "$MODULE_LOG_LEVEL_OVERRIDE" \
  

