package testing

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"

	"github.com/datachainlab/ibc-ethmultisig-client/modules/relay/ethmultisig/wallet"
	"github.com/datachainlab/ibc-ethmultisig-client/pkg/contract/multisigclient"
)

type Chain struct {
	chainID        int64
	mnemonicPhrase string
	keys           map[uint32]*ecdsa.PrivateKey

	ETHClient      *ethclient.Client
	ContractConfig ContractConfig

	ibcHost        ibchost.Ibchost
	multisigClient multisigclient.Multisigclient
}

func NewChain(t *testing.T, rpcAddr string, mnemonicPhrase string, ccfg ContractConfig) *Chain {
	ethc, err := NewETHClient(rpcAddr)
	if err != nil {
		panic(err)
	}
	ibcHost, err := ibchost.NewIbchost(ccfg.GetIBCHostAddress(), ethc)
	if err != nil {
		panic(err)
	}
	msc, err := multisigclient.NewMultisigclient(ccfg.GetMultisigClientAddress(), ethc)
	if err != nil {
		panic(err)
	}
	return &Chain{
		ETHClient:      ethc,
		ContractConfig: ccfg,

		ibcHost:        *ibcHost,
		multisigClient: *msc,

		chainID:        1337,
		mnemonicPhrase: mnemonicPhrase,
		keys:           make(map[uint32]*ecdsa.PrivateKey),
	}
}

func (chain *Chain) TxSync(ctx context.Context, tx *gethtypes.Transaction) error {
	var receipt *gethtypes.Receipt
	err := retry.Do(
		func() error {
			rc, err := chain.ETHClient.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				return err
			}
			receipt = rc
			return nil
		},
		// TODO make these configurable
		retry.Delay(1*time.Second),
		retry.Attempts(10),
	)
	if err != nil {
		return err
	}
	if receipt.Status == gethtypes.ReceiptStatusSuccessful {
		return nil
	} else {
		return fmt.Errorf("failed to call transaction: err='%v' rc='%v'", err, receipt)
	}
}

func (chain *Chain) TxSyncIfNoError(ctx context.Context) func(tx *gethtypes.Transaction, err error) error {
	return func(tx *gethtypes.Transaction, err error) error {
		if err != nil {
			return err
		}
		return chain.TxSync(ctx, tx)
	}
}

func (chain *Chain) TxOpts(ctx context.Context, index uint32) *bind.TransactOpts {
	return makeGenTxOpts(big.NewInt(chain.chainID), chain.prvKey(index))(ctx)
}

func (chain *Chain) CallOpts(ctx context.Context, index uint32) *bind.CallOpts {
	opts := chain.TxOpts(ctx, index)
	return &bind.CallOpts{
		From:    opts.From,
		Context: opts.Context,
	}
}

func (chain *Chain) prvKey(index uint32) *ecdsa.PrivateKey {
	key, ok := chain.keys[index]
	if ok {
		return key
	}
	key, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(chain.mnemonicPhrase, fmt.Sprintf("m/44'/60'/0'/0/%v", index))
	if err != nil {
		panic(err)
	}
	chain.keys[index] = key
	return key
}

type GenTxOpts func(ctx context.Context) *bind.TransactOpts

func makeGenTxOpts(chainID *big.Int, prv *ecdsa.PrivateKey) GenTxOpts {
	signer := gethtypes.NewEIP155Signer(chainID)
	addr := gethcrypto.PubkeyToAddress(prv.PublicKey)
	return func(ctx context.Context) *bind.TransactOpts {
		return &bind.TransactOpts{
			From:     addr,
			GasLimit: 6382056,
			Signer: func(address common.Address, tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
				if address != addr {
					return nil, errors.New("not authorized to sign this account")
				}
				signature, err := gethcrypto.Sign(signer.Hash(tx).Bytes(), prv)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
		}
	}
}

type ContractConfig interface {
	GetIBCHostAddress() common.Address
	GetMultisigClientAddress() common.Address
}

func NewETHClient(rpcAddr string) (*ethclient.Client, error) {
	conn, err := rpc.DialHTTP(rpcAddr)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(conn), nil
}
