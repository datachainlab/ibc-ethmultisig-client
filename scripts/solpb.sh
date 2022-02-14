#!/usr/bin/env bash
set -e

if [ -z "$SOLPB_DIR" ]; then
    echo "variable SOLPB_DIR must be set"
    exit 1
fi

if [ -z "$SOLPB_EXTERNAL_RUNTIME_REPO" ]; then
    echo "variable SOLPB_EXTERNAL_RUNTIME_REPO must be set"
    exit 1
fi

echo ""
echo "Use the runtime packages in yui-ibc-solidity"
echo "package: $SOLPB_EXTERNAL_RUNTIME_REPO"

for file in $(find ./proto/ibc/lightclients -name '*.proto')
do
  echo "Generating "$file
  protoc \
    -I$(pwd)/proto \
    -I "third_party/proto" \
    -I${SOLPB_DIR}/protobuf-solidity/src/protoc/include  \
    --plugin=protoc-gen-sol=${SOLPB_DIR}/protobuf-solidity/src/protoc/plugin/gen_sol.py  \
    --"sol_out=use_runtime=${SOLPB_EXTERNAL_RUNTIME_REPO}ProtoBufRuntime.sol&solc_version=0.8.9&ignore_protos=gogoproto/gogo.proto:$(pwd)" $(pwd)/$file
done

# FIXME delete here after modifying solidity-protobuf
sed -i -E "s#(^import +\")([\.|\/])+(Client.sol\";$)#\1$SOLPB_EXTERNAL_RUNTIME_REPO\3#" ./contracts/core/types/ethmultisig.sol
rm -rf contracts/core/types/Client.sol
