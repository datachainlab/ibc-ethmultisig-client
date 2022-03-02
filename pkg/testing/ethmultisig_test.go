package testing

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	"github.com/stretchr/testify/suite"

	ethmultisigtypes "github.com/datachainlab/ibc-ethmultisig-client/modules/light-clients/xx-ethmultisig/types"
	"github.com/datachainlab/ibc-ethmultisig-client/modules/relay/ethmultisig"
	"github.com/datachainlab/ibc-ethmultisig-client/pkg/consts"
	"github.com/datachainlab/ibc-ethmultisig-client/pkg/contract/multisigclient"
)

const testMnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ETHMultisigTestSuite struct {
	suite.Suite

	chain *Chain
	cdc   codec.ProtoCodecMarshaler
}

func (suite *ETHMultisigTestSuite) SetupTest() {
	suite.chain = NewChain(suite.T(), "http://127.0.0.1:8545", testMnemonicPhrase, consts.Contract)
	registry := codectypes.NewInterfaceRegistry()
	ethmultisigtypes.RegisterInterfaces(registry)
	suite.cdc = codec.NewProtoCodec(registry)
}

func (suite *ETHMultisigTestSuite) TestMultisig() {
	ctx := context.TODO()

	const (
		diversifier          = "tester"
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	proofHeight := clienttypes.NewHeight(0, 1)
	prefix := []byte("ibc")

	prover := ethmultisig.NewETHMultisig(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, prefix)

	consensusState := makeMultisigConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		diversifier,
		uint64(time.Now().UnixNano()),
	)
	anyConsensusStateBytes, err := suite.cdc.MarshalInterface(consensusState)
	suite.Require().NoError(err)
	err = suite.chain.TxSyncIfNoError(ctx)(
		suite.chain.ibcHost.SetConsensusState(
			suite.chain.TxOpts(ctx, 0),
			clientID,
			ibchost.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			anyConsensusStateBytes,
		))
	suite.Require().NoError(err)

	// VerifyClientState
	{
		targetClientState := makeMultisigClientState(1)
		proofClient, _, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
		suite.Require().NoError(err)
		anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofClient)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyClientState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			}, prefix, counterpartyClientID, proofBytes, anyClientStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyClientConsensusState consensusHeight.revisionNumber=0
	{
		targetConsensusState := makeMultisigConsensusState(nil, "tester2", uint64(time.Now().UnixNano()))
		consensusHeight := clienttypes.NewHeight(0, 100)
		proofConsensus, _, err := prover.SignConsensusState(proofHeight, counterpartyClientID, consensusHeight, targetConsensusState)
		suite.Require().NoError(err)
		anyConsensusStateBytes, err := suite.cdc.MarshalInterface(targetConsensusState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofConsensus)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyClientConsensusState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			counterpartyClientID,
			multisigclient.HeightData{
				RevisionNumber: consensusHeight.RevisionNumber,
				RevisionHeight: consensusHeight.RevisionHeight,
			},
			prefix, proofBytes, anyConsensusStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyClientConsensusState consensusHeight.revisionNumber=1
	{
		targetConsensusState := makeMultisigConsensusState(nil, "tester2", uint64(time.Now().UnixNano()))
		consensusHeight := clienttypes.NewHeight(1, 100)
		proofConsensus, _, err := prover.SignConsensusState(proofHeight, counterpartyClientID, consensusHeight, targetConsensusState)
		suite.Require().NoError(err)
		anyConsensusStateBytes, err := suite.cdc.MarshalInterface(targetConsensusState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofConsensus)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyClientConsensusState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			counterpartyClientID,
			multisigclient.HeightData{
				RevisionNumber: consensusHeight.RevisionNumber,
				RevisionHeight: consensusHeight.RevisionHeight,
			},
			prefix, proofBytes, anyConsensusStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyConnectionState
	{
		const connectionID = "connection-0"
		targetConnection := conntypes.NewConnectionEnd(conntypes.INIT, counterpartyClientID, conntypes.NewCounterparty(clientID, connectionID, types.NewMerklePrefix([]byte("ibc"))), []*conntypes.Version{}, 0)
		proof, _, err := prover.SignConnectionState(proofHeight, connectionID, targetConnection)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		connectionBytes, err := suite.cdc.Marshal(&targetConnection)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyConnectionState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			prefix, proofBytes, connectionID, connectionBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	const portID, channelID, cpPortID, cpChannelID = "port-0", "channel-0", "port-1", "channel-1"

	// VerifyChannelstate
	{
		targetChannel := chantypes.NewChannel(chantypes.INIT, chantypes.UNORDERED, chantypes.NewCounterparty(cpPortID, cpChannelID), []string{"connection-0"}, "1")
		proof, _, err := prover.SignChannelState(proofHeight, portID, channelID, targetChannel)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		channelBytes, err := suite.cdc.Marshal(&targetChannel)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyChannelState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			prefix, proofBytes, portID, channelID, channelBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketCommitment
	{
		commitment := sha256.Sum256([]byte("test"))
		proof, _, err := prover.SignPacketState(proofHeight, portID, channelID, 1, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyPacketCommitment(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			0, 0,
			prefix, proofBytes, portID, channelID, 1, commitment,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketAcknowledgement
	{
		acknowledgement := []byte("ack")
		commitment := sha256.Sum256(acknowledgement)
		proof, _, err := prover.SignPacketAcknowledgementState(proofHeight, portID, channelID, 1, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyPacketAcknowledgement(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			0, 0,
			prefix, proofBytes, portID, channelID, 1, acknowledgement,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}
}

func (suite *ETHMultisigTestSuite) TestMultisigSign() {
	const (
		diversifier          = "tester"
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	proofHeight := clienttypes.NewHeight(0, 1)
	prefix := []byte("ibc")

	prover := ethmultisig.NewETHMultisig(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, prefix)

	targetClientState := makeMultisigClientState(1)
	proofClient, signBytes, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
	suite.Require().NoError(err)

	err = ethmultisigtypes.VerifySignature(prover.Addresses(), proofClient, signBytes)
	suite.Require().NoError(err)
}

func makeMultisigClientState(latestHeight uint64) *ethmultisigtypes.ClientState {
	return &ethmultisigtypes.ClientState{
		LatestHeight: client.Height{
			RevisionNumber: 0,
			RevisionHeight: latestHeight,
		},
	}
}

func makeMultisigConsensusState(addresses []common.Address, diversifier string, timestamp uint64) *ethmultisigtypes.ConsensusState {
	var addrs [][]byte
	for _, addr := range addresses {
		addrs = append(addrs, addr[:])
	}
	return &ethmultisigtypes.ConsensusState{
		Addresses:   addrs,
		Diversifier: diversifier,
		Timestamp:   timestamp,
	}
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ETHMultisigTestSuite))
}
