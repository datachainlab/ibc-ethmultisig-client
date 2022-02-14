// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisigclient

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type ClientStateData struct {
	LatestHeight HeightData
	FrozenHeight HeightData
}

// ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateData struct {
	Addresses   [][]byte
	Diversifier string
	Timestamp   uint64
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// MultiSignatureData is an auto generated low-level Go binding around an user-defined struct.
type MultiSignatureData struct {
	Signatures [][]byte
	Timestamp  uint64
}

// MultisigclientABI is the input ABI used to generate the binding from.
const MultisigclientABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getTimestampAtHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"checkHeaderAndUpdateState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"addresses\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structMultiSignature.Data\",\"name\":\"multisig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signBytes\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"consensusHeight\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientConsensusState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"verifyConnectionState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channelBytes\",\"type\":\"bytes\"}],\"name\":\"verifyChannelState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"commitmentBytes\",\"type\":\"bytes32\"}],\"name\":\"verifyPacketCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"verifyPacketAcknowledgement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"latest_height\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"frozen_height\",\"type\":\"tuple\"}],\"internalType\":\"structClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"addresses\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeClientStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"consensusHeight\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeConsensusStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connection\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeConnectionStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channel\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeChannelStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"path\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"packetCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makePacketSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"path\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"packetAcknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makePacketAcknowledgementSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]"

// Multisigclient is an auto generated Go binding around an Ethereum contract.
type Multisigclient struct {
	MultisigclientCaller     // Read-only binding to the contract
	MultisigclientTransactor // Write-only binding to the contract
	MultisigclientFilterer   // Log filterer for contract events
}

// MultisigclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigclientSession struct {
	Contract     *Multisigclient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigclientCallerSession struct {
	Contract *MultisigclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultisigclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigclientTransactorSession struct {
	Contract     *MultisigclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultisigclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigclientRaw struct {
	Contract *Multisigclient // Generic contract binding to access the raw methods on
}

// MultisigclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigclientCallerRaw struct {
	Contract *MultisigclientCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigclientTransactorRaw struct {
	Contract *MultisigclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisigclient creates a new instance of Multisigclient, bound to a specific deployed contract.
func NewMultisigclient(address common.Address, backend bind.ContractBackend) (*Multisigclient, error) {
	contract, err := bindMultisigclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisigclient{MultisigclientCaller: MultisigclientCaller{contract: contract}, MultisigclientTransactor: MultisigclientTransactor{contract: contract}, MultisigclientFilterer: MultisigclientFilterer{contract: contract}}, nil
}

// NewMultisigclientCaller creates a new read-only instance of Multisigclient, bound to a specific deployed contract.
func NewMultisigclientCaller(address common.Address, caller bind.ContractCaller) (*MultisigclientCaller, error) {
	contract, err := bindMultisigclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigclientCaller{contract: contract}, nil
}

// NewMultisigclientTransactor creates a new write-only instance of Multisigclient, bound to a specific deployed contract.
func NewMultisigclientTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigclientTransactor, error) {
	contract, err := bindMultisigclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigclientTransactor{contract: contract}, nil
}

// NewMultisigclientFilterer creates a new log filterer instance of Multisigclient, bound to a specific deployed contract.
func NewMultisigclientFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigclientFilterer, error) {
	contract, err := bindMultisigclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigclientFilterer{contract: contract}, nil
}

// bindMultisigclient binds a generic wrapper to an already deployed contract.
func bindMultisigclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisigclient *MultisigclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisigclient.Contract.MultisigclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisigclient *MultisigclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigclient.Contract.MultisigclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisigclient *MultisigclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisigclient.Contract.MultisigclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisigclient *MultisigclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisigclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisigclient *MultisigclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisigclient *MultisigclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisigclient.Contract.contract.Transact(opts, method, params...)
}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xfe4391fd.
//
// Solidity: function checkHeaderAndUpdateState(address , string , bytes , bytes ) view returns(bytes, bytes, (uint64,uint64))
func (_Multisigclient *MultisigclientCaller) CheckHeaderAndUpdateState(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 []byte, arg3 []byte) ([]byte, []byte, HeightData, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "checkHeaderAndUpdateState", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([]byte), *new([]byte), *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new(HeightData)).(*HeightData)

	return out0, out1, out2, err

}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xfe4391fd.
//
// Solidity: function checkHeaderAndUpdateState(address , string , bytes , bytes ) view returns(bytes, bytes, (uint64,uint64))
func (_Multisigclient *MultisigclientSession) CheckHeaderAndUpdateState(arg0 common.Address, arg1 string, arg2 []byte, arg3 []byte) ([]byte, []byte, HeightData, error) {
	return _Multisigclient.Contract.CheckHeaderAndUpdateState(&_Multisigclient.CallOpts, arg0, arg1, arg2, arg3)
}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xfe4391fd.
//
// Solidity: function checkHeaderAndUpdateState(address , string , bytes , bytes ) view returns(bytes, bytes, (uint64,uint64))
func (_Multisigclient *MultisigclientCallerSession) CheckHeaderAndUpdateState(arg0 common.Address, arg1 string, arg2 []byte, arg3 []byte) ([]byte, []byte, HeightData, error) {
	return _Multisigclient.Contract.CheckHeaderAndUpdateState(&_Multisigclient.CallOpts, arg0, arg1, arg2, arg3)
}

// GetClientState is a free data retrieval call binding the contract method 0xff33b3e6.
//
// Solidity: function getClientState(address host, string clientId) view returns(((uint64,uint64),(uint64,uint64)) clientState, bool found)
func (_Multisigclient *MultisigclientCaller) GetClientState(opts *bind.CallOpts, host common.Address, clientId string) (struct {
	ClientState ClientStateData
	Found       bool
}, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "getClientState", host, clientId)

	outstruct := new(struct {
		ClientState ClientStateData
		Found       bool
	})

	outstruct.ClientState = out[0].(ClientStateData)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// GetClientState is a free data retrieval call binding the contract method 0xff33b3e6.
//
// Solidity: function getClientState(address host, string clientId) view returns(((uint64,uint64),(uint64,uint64)) clientState, bool found)
func (_Multisigclient *MultisigclientSession) GetClientState(host common.Address, clientId string) (struct {
	ClientState ClientStateData
	Found       bool
}, error) {
	return _Multisigclient.Contract.GetClientState(&_Multisigclient.CallOpts, host, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0xff33b3e6.
//
// Solidity: function getClientState(address host, string clientId) view returns(((uint64,uint64),(uint64,uint64)) clientState, bool found)
func (_Multisigclient *MultisigclientCallerSession) GetClientState(host common.Address, clientId string) (struct {
	ClientState ClientStateData
	Found       bool
}, error) {
	return _Multisigclient.Contract.GetClientState(&_Multisigclient.CallOpts, host, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xdd8dd65f.
//
// Solidity: function getConsensusState(address host, string clientId, (uint64,uint64) height) view returns((bytes[],string,uint64) consensusState, bool found)
func (_Multisigclient *MultisigclientCaller) GetConsensusState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData) (struct {
	ConsensusState ConsensusStateData
	Found          bool
}, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "getConsensusState", host, clientId, height)

	outstruct := new(struct {
		ConsensusState ConsensusStateData
		Found          bool
	})

	outstruct.ConsensusState = out[0].(ConsensusStateData)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xdd8dd65f.
//
// Solidity: function getConsensusState(address host, string clientId, (uint64,uint64) height) view returns((bytes[],string,uint64) consensusState, bool found)
func (_Multisigclient *MultisigclientSession) GetConsensusState(host common.Address, clientId string, height HeightData) (struct {
	ConsensusState ConsensusStateData
	Found          bool
}, error) {
	return _Multisigclient.Contract.GetConsensusState(&_Multisigclient.CallOpts, host, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xdd8dd65f.
//
// Solidity: function getConsensusState(address host, string clientId, (uint64,uint64) height) view returns((bytes[],string,uint64) consensusState, bool found)
func (_Multisigclient *MultisigclientCallerSession) GetConsensusState(host common.Address, clientId string, height HeightData) (struct {
	ConsensusState ConsensusStateData
	Found          bool
}, error) {
	return _Multisigclient.Contract.GetConsensusState(&_Multisigclient.CallOpts, host, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0xfe017750.
//
// Solidity: function getLatestHeight(address host, string clientId) view returns((uint64,uint64), bool)
func (_Multisigclient *MultisigclientCaller) GetLatestHeight(opts *bind.CallOpts, host common.Address, clientId string) (HeightData, bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "getLatestHeight", host, clientId)

	if err != nil {
		return *new(HeightData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0xfe017750.
//
// Solidity: function getLatestHeight(address host, string clientId) view returns((uint64,uint64), bool)
func (_Multisigclient *MultisigclientSession) GetLatestHeight(host common.Address, clientId string) (HeightData, bool, error) {
	return _Multisigclient.Contract.GetLatestHeight(&_Multisigclient.CallOpts, host, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0xfe017750.
//
// Solidity: function getLatestHeight(address host, string clientId) view returns((uint64,uint64), bool)
func (_Multisigclient *MultisigclientCallerSession) GetLatestHeight(host common.Address, clientId string) (HeightData, bool, error) {
	return _Multisigclient.Contract.GetLatestHeight(&_Multisigclient.CallOpts, host, clientId)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xa65a0aba.
//
// Solidity: function getTimestampAtHeight(address host, string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Multisigclient *MultisigclientCaller) GetTimestampAtHeight(opts *bind.CallOpts, host common.Address, clientId string, height HeightData) (uint64, bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "getTimestampAtHeight", host, clientId, height)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xa65a0aba.
//
// Solidity: function getTimestampAtHeight(address host, string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Multisigclient *MultisigclientSession) GetTimestampAtHeight(host common.Address, clientId string, height HeightData) (uint64, bool, error) {
	return _Multisigclient.Contract.GetTimestampAtHeight(&_Multisigclient.CallOpts, host, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xa65a0aba.
//
// Solidity: function getTimestampAtHeight(address host, string clientId, (uint64,uint64) height) view returns(uint64, bool)
func (_Multisigclient *MultisigclientCallerSession) GetTimestampAtHeight(host common.Address, clientId string, height HeightData) (uint64, bool, error) {
	return _Multisigclient.Contract.GetTimestampAtHeight(&_Multisigclient.CallOpts, host, clientId, height)
}

// MakeChannelStateSignBytes is a free data retrieval call binding the contract method 0x1b380851.
//
// Solidity: function makeChannelStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string portID, string channelID, bytes channel, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCaller) MakeChannelStateSignBytes(opts *bind.CallOpts, height HeightData, timestamp uint64, diversifier string, portID string, channelID string, channel []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "makeChannelStateSignBytes", height, timestamp, diversifier, portID, channelID, channel, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeChannelStateSignBytes is a free data retrieval call binding the contract method 0x1b380851.
//
// Solidity: function makeChannelStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string portID, string channelID, bytes channel, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientSession) MakeChannelStateSignBytes(height HeightData, timestamp uint64, diversifier string, portID string, channelID string, channel []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeChannelStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, portID, channelID, channel, prefix)
}

// MakeChannelStateSignBytes is a free data retrieval call binding the contract method 0x1b380851.
//
// Solidity: function makeChannelStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string portID, string channelID, bytes channel, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCallerSession) MakeChannelStateSignBytes(height HeightData, timestamp uint64, diversifier string, portID string, channelID string, channel []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeChannelStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, portID, channelID, channel, prefix)
}

// MakeClientStateSignBytes is a free data retrieval call binding the contract method 0xb000a41d.
//
// Solidity: function makeClientStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string clientID, bytes clientState, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCaller) MakeClientStateSignBytes(opts *bind.CallOpts, height HeightData, timestamp uint64, diversifier string, clientID string, clientState []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "makeClientStateSignBytes", height, timestamp, diversifier, clientID, clientState, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeClientStateSignBytes is a free data retrieval call binding the contract method 0xb000a41d.
//
// Solidity: function makeClientStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string clientID, bytes clientState, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientSession) MakeClientStateSignBytes(height HeightData, timestamp uint64, diversifier string, clientID string, clientState []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeClientStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, clientID, clientState, prefix)
}

// MakeClientStateSignBytes is a free data retrieval call binding the contract method 0xb000a41d.
//
// Solidity: function makeClientStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string clientID, bytes clientState, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCallerSession) MakeClientStateSignBytes(height HeightData, timestamp uint64, diversifier string, clientID string, clientState []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeClientStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, clientID, clientState, prefix)
}

// MakeConnectionStateSignBytes is a free data retrieval call binding the contract method 0x291d7cc7.
//
// Solidity: function makeConnectionStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string connectionID, bytes connection, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCaller) MakeConnectionStateSignBytes(opts *bind.CallOpts, height HeightData, timestamp uint64, diversifier string, connectionID string, connection []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "makeConnectionStateSignBytes", height, timestamp, diversifier, connectionID, connection, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeConnectionStateSignBytes is a free data retrieval call binding the contract method 0x291d7cc7.
//
// Solidity: function makeConnectionStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string connectionID, bytes connection, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientSession) MakeConnectionStateSignBytes(height HeightData, timestamp uint64, diversifier string, connectionID string, connection []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeConnectionStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, connectionID, connection, prefix)
}

// MakeConnectionStateSignBytes is a free data retrieval call binding the contract method 0x291d7cc7.
//
// Solidity: function makeConnectionStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string connectionID, bytes connection, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCallerSession) MakeConnectionStateSignBytes(height HeightData, timestamp uint64, diversifier string, connectionID string, connection []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeConnectionStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, connectionID, connection, prefix)
}

// MakeConsensusStateSignBytes is a free data retrieval call binding the contract method 0x8c5d60b4.
//
// Solidity: function makeConsensusStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string clientID, (uint64,uint64) consensusHeight, bytes consensusState, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCaller) MakeConsensusStateSignBytes(opts *bind.CallOpts, height HeightData, timestamp uint64, diversifier string, clientID string, consensusHeight HeightData, consensusState []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "makeConsensusStateSignBytes", height, timestamp, diversifier, clientID, consensusHeight, consensusState, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeConsensusStateSignBytes is a free data retrieval call binding the contract method 0x8c5d60b4.
//
// Solidity: function makeConsensusStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string clientID, (uint64,uint64) consensusHeight, bytes consensusState, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientSession) MakeConsensusStateSignBytes(height HeightData, timestamp uint64, diversifier string, clientID string, consensusHeight HeightData, consensusState []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeConsensusStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, clientID, consensusHeight, consensusState, prefix)
}

// MakeConsensusStateSignBytes is a free data retrieval call binding the contract method 0x8c5d60b4.
//
// Solidity: function makeConsensusStateSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, string clientID, (uint64,uint64) consensusHeight, bytes consensusState, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCallerSession) MakeConsensusStateSignBytes(height HeightData, timestamp uint64, diversifier string, clientID string, consensusHeight HeightData, consensusState []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakeConsensusStateSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, clientID, consensusHeight, consensusState, prefix)
}

// MakePacketAcknowledgementSignBytes is a free data retrieval call binding the contract method 0x9e4d202b.
//
// Solidity: function makePacketAcknowledgementSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, bytes32 path, bytes packetAcknowledgement, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCaller) MakePacketAcknowledgementSignBytes(opts *bind.CallOpts, height HeightData, timestamp uint64, diversifier string, path [32]byte, packetAcknowledgement []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "makePacketAcknowledgementSignBytes", height, timestamp, diversifier, path, packetAcknowledgement, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakePacketAcknowledgementSignBytes is a free data retrieval call binding the contract method 0x9e4d202b.
//
// Solidity: function makePacketAcknowledgementSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, bytes32 path, bytes packetAcknowledgement, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientSession) MakePacketAcknowledgementSignBytes(height HeightData, timestamp uint64, diversifier string, path [32]byte, packetAcknowledgement []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakePacketAcknowledgementSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, path, packetAcknowledgement, prefix)
}

// MakePacketAcknowledgementSignBytes is a free data retrieval call binding the contract method 0x9e4d202b.
//
// Solidity: function makePacketAcknowledgementSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, bytes32 path, bytes packetAcknowledgement, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCallerSession) MakePacketAcknowledgementSignBytes(height HeightData, timestamp uint64, diversifier string, path [32]byte, packetAcknowledgement []byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakePacketAcknowledgementSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, path, packetAcknowledgement, prefix)
}

// MakePacketSignBytes is a free data retrieval call binding the contract method 0xba32f992.
//
// Solidity: function makePacketSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, bytes32 path, bytes32 packetCommitment, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCaller) MakePacketSignBytes(opts *bind.CallOpts, height HeightData, timestamp uint64, diversifier string, path [32]byte, packetCommitment [32]byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "makePacketSignBytes", height, timestamp, diversifier, path, packetCommitment, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakePacketSignBytes is a free data retrieval call binding the contract method 0xba32f992.
//
// Solidity: function makePacketSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, bytes32 path, bytes32 packetCommitment, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientSession) MakePacketSignBytes(height HeightData, timestamp uint64, diversifier string, path [32]byte, packetCommitment [32]byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakePacketSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, path, packetCommitment, prefix)
}

// MakePacketSignBytes is a free data retrieval call binding the contract method 0xba32f992.
//
// Solidity: function makePacketSignBytes((uint64,uint64) height, uint64 timestamp, string diversifier, bytes32 path, bytes32 packetCommitment, bytes prefix) pure returns(bytes)
func (_Multisigclient *MultisigclientCallerSession) MakePacketSignBytes(height HeightData, timestamp uint64, diversifier string, path [32]byte, packetCommitment [32]byte, prefix []byte) ([]byte, error) {
	return _Multisigclient.Contract.MakePacketSignBytes(&_Multisigclient.CallOpts, height, timestamp, diversifier, path, packetCommitment, prefix)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0xc9f13d26.
//
// Solidity: function verifyChannelState(address host, string clientId, (uint64,uint64) height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifyChannelState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifyChannelState", host, clientId, height, prefix, proof, portId, channelId, channelBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyChannelState is a free data retrieval call binding the contract method 0xc9f13d26.
//
// Solidity: function verifyChannelState(address host, string clientId, (uint64,uint64) height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Multisigclient *MultisigclientSession) VerifyChannelState(host common.Address, clientId string, height HeightData, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyChannelState(&_Multisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, channelBytes)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0xc9f13d26.
//
// Solidity: function verifyChannelState(address host, string clientId, (uint64,uint64) height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifyChannelState(host common.Address, clientId string, height HeightData, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyChannelState(&_Multisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, channelBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0x8eaa4a38.
//
// Solidity: function verifyClientConsensusState(address host, string clientId, (uint64,uint64) height, string counterpartyClientIdentifier, (uint64,uint64) consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifyClientConsensusState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData, counterpartyClientIdentifier string, consensusHeight HeightData, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifyClientConsensusState", host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0x8eaa4a38.
//
// Solidity: function verifyClientConsensusState(address host, string clientId, (uint64,uint64) height, string counterpartyClientIdentifier, (uint64,uint64) consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Multisigclient *MultisigclientSession) VerifyClientConsensusState(host common.Address, clientId string, height HeightData, counterpartyClientIdentifier string, consensusHeight HeightData, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyClientConsensusState(&_Multisigclient.CallOpts, host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0x8eaa4a38.
//
// Solidity: function verifyClientConsensusState(address host, string clientId, (uint64,uint64) height, string counterpartyClientIdentifier, (uint64,uint64) consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifyClientConsensusState(host common.Address, clientId string, height HeightData, counterpartyClientIdentifier string, consensusHeight HeightData, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyClientConsensusState(&_Multisigclient.CallOpts, host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0x9e6671bd.
//
// Solidity: function verifyClientState(address host, string clientId, (uint64,uint64) height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifyClientState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifyClientState", host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientState is a free data retrieval call binding the contract method 0x9e6671bd.
//
// Solidity: function verifyClientState(address host, string clientId, (uint64,uint64) height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Multisigclient *MultisigclientSession) VerifyClientState(host common.Address, clientId string, height HeightData, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyClientState(&_Multisigclient.CallOpts, host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0x9e6671bd.
//
// Solidity: function verifyClientState(address host, string clientId, (uint64,uint64) height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifyClientState(host common.Address, clientId string, height HeightData, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyClientState(&_Multisigclient.CallOpts, host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xa8432fa1.
//
// Solidity: function verifyConnectionState(address host, string clientId, (uint64,uint64) height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifyConnectionState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifyConnectionState", host, clientId, height, prefix, proof, connectionId, connectionBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xa8432fa1.
//
// Solidity: function verifyConnectionState(address host, string clientId, (uint64,uint64) height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Multisigclient *MultisigclientSession) VerifyConnectionState(host common.Address, clientId string, height HeightData, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyConnectionState(&_Multisigclient.CallOpts, host, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xa8432fa1.
//
// Solidity: function verifyConnectionState(address host, string clientId, (uint64,uint64) height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifyConnectionState(host common.Address, clientId string, height HeightData, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyConnectionState(&_Multisigclient.CallOpts, host, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x54596a36.
//
// Solidity: function verifyPacketAcknowledgement(address host, string clientId, (uint64,uint64) height, uint64 , uint64 , bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes acknowledgement) view returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifyPacketAcknowledgement(opts *bind.CallOpts, host common.Address, clientId string, height HeightData, arg3 uint64, arg4 uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, acknowledgement []byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifyPacketAcknowledgement", host, clientId, height, arg3, arg4, prefix, proof, portId, channelId, sequence, acknowledgement)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x54596a36.
//
// Solidity: function verifyPacketAcknowledgement(address host, string clientId, (uint64,uint64) height, uint64 , uint64 , bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes acknowledgement) view returns(bool)
func (_Multisigclient *MultisigclientSession) VerifyPacketAcknowledgement(host common.Address, clientId string, height HeightData, arg3 uint64, arg4 uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, acknowledgement []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyPacketAcknowledgement(&_Multisigclient.CallOpts, host, clientId, height, arg3, arg4, prefix, proof, portId, channelId, sequence, acknowledgement)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x54596a36.
//
// Solidity: function verifyPacketAcknowledgement(address host, string clientId, (uint64,uint64) height, uint64 , uint64 , bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes acknowledgement) view returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifyPacketAcknowledgement(host common.Address, clientId string, height HeightData, arg3 uint64, arg4 uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, acknowledgement []byte) (bool, error) {
	return _Multisigclient.Contract.VerifyPacketAcknowledgement(&_Multisigclient.CallOpts, host, clientId, height, arg3, arg4, prefix, proof, portId, channelId, sequence, acknowledgement)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0xa0aaade7.
//
// Solidity: function verifyPacketCommitment(address host, string clientId, (uint64,uint64) height, uint64 , uint64 , bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifyPacketCommitment(opts *bind.CallOpts, host common.Address, clientId string, height HeightData, arg3 uint64, arg4 uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifyPacketCommitment", host, clientId, height, arg3, arg4, prefix, proof, portId, channelId, sequence, commitmentBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0xa0aaade7.
//
// Solidity: function verifyPacketCommitment(address host, string clientId, (uint64,uint64) height, uint64 , uint64 , bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Multisigclient *MultisigclientSession) VerifyPacketCommitment(host common.Address, clientId string, height HeightData, arg3 uint64, arg4 uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Multisigclient.Contract.VerifyPacketCommitment(&_Multisigclient.CallOpts, host, clientId, height, arg3, arg4, prefix, proof, portId, channelId, sequence, commitmentBytes)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0xa0aaade7.
//
// Solidity: function verifyPacketCommitment(address host, string clientId, (uint64,uint64) height, uint64 , uint64 , bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifyPacketCommitment(host common.Address, clientId string, height HeightData, arg3 uint64, arg4 uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Multisigclient.Contract.VerifyPacketCommitment(&_Multisigclient.CallOpts, host, clientId, height, arg3, arg4, prefix, proof, portId, channelId, sequence, commitmentBytes)
}

// VerifySignature is a free data retrieval call binding the contract method 0xe332d75c.
//
// Solidity: function verifySignature((bytes[],string,uint64) consensusState, (bytes[],uint64) multisig, bytes signBytes) pure returns(bool)
func (_Multisigclient *MultisigclientCaller) VerifySignature(opts *bind.CallOpts, consensusState ConsensusStateData, multisig MultiSignatureData, signBytes []byte) (bool, error) {
	var out []interface{}
	err := _Multisigclient.contract.Call(opts, &out, "verifySignature", consensusState, multisig, signBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0xe332d75c.
//
// Solidity: function verifySignature((bytes[],string,uint64) consensusState, (bytes[],uint64) multisig, bytes signBytes) pure returns(bool)
func (_Multisigclient *MultisigclientSession) VerifySignature(consensusState ConsensusStateData, multisig MultiSignatureData, signBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifySignature(&_Multisigclient.CallOpts, consensusState, multisig, signBytes)
}

// VerifySignature is a free data retrieval call binding the contract method 0xe332d75c.
//
// Solidity: function verifySignature((bytes[],string,uint64) consensusState, (bytes[],uint64) multisig, bytes signBytes) pure returns(bool)
func (_Multisigclient *MultisigclientCallerSession) VerifySignature(consensusState ConsensusStateData, multisig MultiSignatureData, signBytes []byte) (bool, error) {
	return _Multisigclient.Contract.VerifySignature(&_Multisigclient.CallOpts, consensusState, multisig, signBytes)
}
