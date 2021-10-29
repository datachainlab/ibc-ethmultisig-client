package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "<%= MultisigClientAddress; %>"
	IBCHostAddress = "<%= IBCHostAddress; %>"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetMultisigClientAddress() common.Address {
	return common.HexToAddress(MultisigClientAddress)
}

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}
