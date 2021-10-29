package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "0x361552A65C96621003C62C5971b910a1fdC9ba78"
	IBCHostAddress = "0x9eBF3956EE45B2b9F1fC85FB8990ce6be52F47a6"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetMultisigClientAddress() common.Address {
	return common.HexToAddress(MultisigClientAddress)
}

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}
