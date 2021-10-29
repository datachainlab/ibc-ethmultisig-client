package types

import (
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"
)

func (cs *ConsensusState) ClientType() string {
	return ClientType
}

// GetRoot returns the commitment root of the consensus state,
// which is used for key-value pair verification.
func (cs *ConsensusState) GetRoot() exported.Root {
	return nil
}

// GetTimestamp returns the timestamp (in nanoseconds) of the consensus state
func (cs *ConsensusState) GetTimestamp() uint64 {
	return cs.Timestamp
}

func (cs *ConsensusState) ValidateBasic() error {
	return nil
}

func (cs *ConsensusState) GetAddresses() []common.Address {
	var addrs []common.Address
	for _, addr := range cs.Addresses {
		addrs = append(addrs, common.BytesToAddress(addr))
	}
	return addrs
}
