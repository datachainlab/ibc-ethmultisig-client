package ethmultisig

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/hyperledger-labs/yui-relayer/core"

	ethmultisigtypes "github.com/datachainlab/ibc-ethmultisig-client/modules/light-clients/xx-ethmultisig/types"
)

// RegisterInterfaces register the module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	ethmultisigtypes.RegisterInterfaces(registry)
	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
}
