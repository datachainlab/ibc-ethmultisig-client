package ethmultisig

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"

	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger-labs/yui-relayer/core"

	ethmultisigclient "github.com/datachainlab/ibc-ethmultisig-client/modules/light-clients/xx-ethmultisig/types"
	"github.com/datachainlab/ibc-ethmultisig-client/modules/relay/ethmultisig/wallet"
)

var _ core.ProverConfigI = (*ProverConfig)(nil)

func (pr ProverConfig) Build(chain core.ChainI) (core.ProverI, error) {
	return NewProver(pr, chain)
}

type Prover struct {
	chain core.ChainI

	diversifier string
	multisig    ETHMultisig
}

var _ core.ProverI = (*Prover)(nil)

func NewProver(pr ProverConfig, chain core.ChainI) (*Prover, error) {
	if len(pr.Wallets) == 0 {
		return nil, fmt.Errorf("at least one wallet is needed")
	}
	var keys []*ecdsa.PrivateKey
	for _, w := range pr.Wallets {
		prv, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(w.Mnemonic, w.HdwPath)
		if err != nil {
			return nil, err
		}
		keys = append(keys, prv)
	}
	multisig := NewETHMultisig(chain.Codec(), pr.Diversifier, keys, []byte(pr.Prefix))
	return &Prover{chain: chain, diversifier: pr.Diversifier, multisig: multisig}, nil
}

// GetChainID returns the chain ID
func (pr *Prover) GetChainID() string {
	return pr.chain.ChainID()
}

// QueryLatestHeader returns the latest header from the chain
func (pr *Prover) QueryLatestHeader() (out core.HeaderI, err error) {
	return &ethmultisigclient.Header{Height: client.Height{RevisionNumber: 0, RevisionHeight: 1}}, nil
}

// GetLatestLightHeight returns the latest height on the light client
func (pr *Prover) GetLatestLightHeight() (int64, error) {
	return 1, nil
}

// CreateMsgCreateClient creates a CreateClientMsg to this chain
func (pr *Prover) CreateMsgCreateClient(clientID string, dstHeader core.HeaderI, signer sdk.AccAddress) (*clienttypes.MsgCreateClient, error) {
	var addresses [][]byte
	for _, addr := range pr.multisig.Addresses() {
		addresses = append(addresses, addr.Bytes())
	}
	clientState := &ethmultisigclient.ClientState{
		LatestHeight: client.Height{RevisionNumber: 0, RevisionHeight: 1},
	}
	consensusState := &ethmultisigclient.ConsensusState{
		Addresses:   addresses,
		Diversifier: pr.diversifier,
		Timestamp:   uint64(time.Now().UnixNano()),
	}
	return clienttypes.NewMsgCreateClient(clientState, consensusState, signer.String())
}

// SetupHeader creates a new header based on a given header
func (pr *Prover) SetupHeader(dst core.LightClientIBCQueryierI, baseSrcHeader core.HeaderI) (core.HeaderI, error) {
	return nil, nil
}

// UpdateLightWithHeader updates a header on the light client and returns the header and height corresponding to the chain
func (pr *Prover) UpdateLightWithHeader() (header core.HeaderI, provableHeight int64, queryableHeight int64, err error) {
	h, err := pr.QueryLatestHeader()
	if err != nil {
		return nil, 0, 0, err
	}
	return h, 0, 0, nil
}

/* Query functions: Prover queries the latest state of the chain and signs it */

// QueryClientConsensusState returns the ClientConsensusState and its proof
func (pr *Prover) QueryClientConsensusStateWithProof(_ int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	h, err := pr.chain.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	res, err := pr.chain.QueryClientConsensusState(h, dstClientConsHeight)
	if err != nil {
		return nil, err
	}
	return pr.SignConsensusStateResponse(res, pr.chain.Path().ClientID, dstClientConsHeight)
}

// QueryClientStateWithProof returns the ClientState and its proof
func (pr *Prover) QueryClientStateWithProof(_ int64) (*clienttypes.QueryClientStateResponse, error) {
	h, err := pr.chain.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	res, err := pr.chain.QueryClientState(h)
	if err != nil {
		return nil, err
	}
	return pr.SignClientStateResponse(res, pr.chain.Path().ClientID)
}

// QueryConnectionWithProof returns the Connection and its proof
func (pr *Prover) QueryConnectionWithProof(_ int64) (*conntypes.QueryConnectionResponse, error) {
	h, err := pr.chain.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	res, err := pr.chain.QueryConnection(h)
	if err != nil {
		return nil, err
	}
	return pr.SignConnectionStateResponse(res, pr.chain.Path().ConnectionID)
}

// QueryChannelWithProof returns the Channel and its proof
func (pr *Prover) QueryChannelWithProof(_ int64) (*chantypes.QueryChannelResponse, error) {
	h, err := pr.chain.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	res, err := pr.chain.QueryChannel(h)
	if err != nil {
		return nil, err
	}
	return pr.SignChannelStateResponse(res, pr.chain.Path().PortID, pr.chain.Path().ChannelID)
}

// QueryPacketCommitmentWithProof returns the packet commitment and its proof
func (pr *Prover) QueryPacketCommitmentWithProof(_ int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	h, err := pr.chain.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	res, err := pr.chain.QueryPacketCommitment(h, seq)
	if err != nil {
		return nil, err
	}
	return pr.SignPacketStateResponse(res, pr.chain.Path().PortID, pr.chain.Path().ChannelID, seq)
}

// QueryPacketAcknowledgementCommitmentWithProof returns the packet acknowledgement commitment and its proof
func (pr *Prover) QueryPacketAcknowledgementCommitmentWithProof(_ int64, seq uint64) (*chantypes.QueryPacketAcknowledgementResponse, error) {
	h, err := pr.chain.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	res, err := pr.chain.QueryPacketAcknowledgementCommitment(h, seq)
	if err != nil {
		return nil, err
	}
	return pr.SignAcknowledgementStateResponse(res, pr.chain.Path().PortID, pr.chain.Path().ChannelID, seq)
}

func (pr *Prover) GetHeight() (clienttypes.Height, error) {
	seq, err := pr.GetSequeunce()
	if err != nil {
		return clienttypes.Height{}, err
	}
	return clienttypes.NewHeight(0, seq), nil
}

// TODO load a sequence value from the persisted storage
func (pr *Prover) GetSequeunce() (uint64, error) {
	return 1, nil
}

func (pr *Prover) SignClientStateResponse(res *clienttypes.QueryClientStateResponse, clientID string) (*clienttypes.QueryClientStateResponse, error) {
	clientState, err := clienttypes.UnpackClientState(res.ClientState)
	if err != nil {
		return nil, err
	}
	res.ProofHeight, err = pr.GetHeight()
	if err != nil {
		return nil, err
	}
	pr.xxxInit(pr.chain.Codec())
	res.Proof, err = marshalProofIfNoError(pr.multisig.SignClientState(res.ProofHeight, clientID, clientState))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pr *Prover) SignConsensusStateResponse(res *clienttypes.QueryConsensusStateResponse, clientID string, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	consensusState, err := clienttypes.UnpackConsensusState(res.ConsensusState)
	if err != nil {
		return nil, err
	}
	res.ProofHeight, err = pr.GetHeight()
	if err != nil {
		return nil, err
	}
	pr.xxxInit(pr.chain.Codec())
	res.Proof, err = marshalProofIfNoError(pr.multisig.SignConsensusState(res.ProofHeight, clientID, dstClientConsHeight, consensusState))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pr *Prover) SignConnectionStateResponse(res *conntypes.QueryConnectionResponse, connectionID string) (*conntypes.QueryConnectionResponse, error) {
	var err error
	res.ProofHeight, err = pr.GetHeight()
	if err != nil {
		return nil, err
	}
	pr.xxxInit(pr.chain.Codec())
	res.Proof, err = marshalProofIfNoError(pr.multisig.SignConnectionState(res.ProofHeight, connectionID, *res.Connection))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pr *Prover) SignChannelStateResponse(res *chantypes.QueryChannelResponse, portID, channelID string) (*chantypes.QueryChannelResponse, error) {
	var err error
	res.ProofHeight, err = pr.GetHeight()
	if err != nil {
		return nil, err
	}
	pr.xxxInit(pr.chain.Codec())
	res.Proof, err = marshalProofIfNoError(pr.multisig.SignChannelState(res.ProofHeight, portID, channelID, *res.Channel))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pr *Prover) SignPacketStateResponse(res *chantypes.QueryPacketCommitmentResponse, portID, channelID string, seq uint64) (*chantypes.QueryPacketCommitmentResponse, error) {
	var err error
	res.ProofHeight, err = pr.GetHeight()
	if err != nil {
		return nil, err
	}
	pr.xxxInit(pr.chain.Codec())
	res.Proof, err = marshalProofIfNoError(pr.multisig.SignPacketState(res.ProofHeight, portID, channelID, seq, res.Commitment))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pr *Prover) SignAcknowledgementStateResponse(res *chantypes.QueryPacketAcknowledgementResponse, portID, channelID string, seq uint64) (*chantypes.QueryPacketAcknowledgementResponse, error) {
	var err error
	res.ProofHeight, err = pr.GetHeight()
	if err != nil {
		return nil, err
	}
	pr.xxxInit(pr.chain.Codec())
	res.Proof, err = marshalProofIfNoError(pr.multisig.SignPacketAcknowledgementState(res.ProofHeight, portID, channelID, seq, res.Acknowledgement))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// xxxInit initializes the codec of ethmultisig
// TODO: This method should be removed after the problem with the prover not giving a codec is fixed
func (pr *Prover) xxxInit(cdc codec.ProtoCodecMarshaler) {
	pr.multisig.cdc = cdc
}

func marshalProofIfNoError(proof *ethmultisigclient.MultiSignature, _ []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	return proto.Marshal(proof)
}
