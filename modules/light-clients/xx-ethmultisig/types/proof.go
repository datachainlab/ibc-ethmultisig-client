package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
)

func VerifySignature(addresses []common.Address, multiSig *MultiSignature, signBytes []byte) error {
	if len(addresses) != len(multiSig.Signatures) {
		return ErrInvalidSignatureCount
	}
	h := crypto.Keccak256(signBytes)
	for i, sig := range multiSig.Signatures {
		if len(sig) != crypto.SignatureLength {
			return fmt.Errorf("signature must be 65 bytes long")
		}
		rpk, err := crypto.SigToPub(h, sig)
		if err != nil {
			return err
		}
		signer := crypto.PubkeyToAddress(*rpk)
		if addresses[i] != signer {
			return fmt.Errorf("signature does not match signer: %v != %v (hash=%v)", addresses[i], signer, h)
		}
	}
	return nil
}

// ClientStateSignBytes returns the sign bytes for verification of the
// client state.
func ClientStateSignBytes(
	cdc codec.BinaryCodec,
	height clienttypes.Height, timestamp uint64,
	diversifier string,
	path []byte,
	clientState exported.ClientState,
) ([]byte, error) {
	clientStateBz, err := cdc.MarshalInterface(clientState)
	if err != nil {
		return nil, err
	}
	data := StateData{
		Path:  path,
		Value: clientStateBz,
	}
	dataBz, err := cdc.Marshal(&data)
	if err != nil {
		return nil, err
	}
	signBytes := &SignBytes{
		Height:      client.Height(height),
		Timestamp:   timestamp,
		Diversifier: diversifier,
		DataType:    CLIENT,
		Data:        dataBz,
	}
	return cdc.Marshal(signBytes)
}

// ConsensusStateSignBytes returns the sign bytes for verification of the
// consensus state.
func ConsensusStateSignBytes(
	cdc codec.BinaryCodec,
	height clienttypes.Height, timestamp uint64,
	diversifier string,
	path []byte,
	consensusState exported.ConsensusState,
) ([]byte, error) {
	consensusStateBz, err := cdc.MarshalInterface(consensusState)
	if err != nil {
		return nil, err
	}
	data := StateData{
		Path:  path,
		Value: consensusStateBz,
	}
	dataBz, err := cdc.Marshal(&data)
	if err != nil {
		return nil, err
	}
	signBytes := &SignBytes{
		Height:      client.Height(height),
		Timestamp:   timestamp,
		Diversifier: diversifier,
		DataType:    CONSENSUS,
		Data:        dataBz,
	}
	return cdc.Marshal(signBytes)
}
