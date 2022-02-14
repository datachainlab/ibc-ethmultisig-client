module github.com/datachainlab/ibc-ethmultisig-client

go 1.16

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/confio/ics23/go v0.6.6
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/solidity-protobuf/protobuf-solidity/src/protoc/go v0.0.0-20220114022221-e2f13f13eba5
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20220214080515-0f917e10509b
	github.com/hyperledger-labs/yui-relayer v0.1.1-0.20210818033701-ef1f6d422958
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef
)

replace (
	github.com/cosmos/ibc-go => github.com/datachainlab/ibc-go v0.0.0-20210623043207-6582d8c965f8
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
