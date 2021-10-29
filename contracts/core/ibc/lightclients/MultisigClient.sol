pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import {IClient} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IClient.sol";
import {IBCHost} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHost.sol";
import {IBCIdentifier} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCIdentifier.sol";
import {GoogleProtobufAny as Any} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/types/GoogleProtobufAny.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/lib/ECRecovery.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/lib/Bytes.sol";
import {
    ClientState,
    ConsensusState,
    Header,
    MultiSignature,
    SignBytes,
    StateData,
    Height
} from "../../types/ethmultisig.sol";

// MultisigClient is a dialect of https://github.com/datachainlab/ibc-multisig-client
contract MultisigClient is IClient {
    using Bytes for bytes;

    struct protoTypes {
        bytes32 clientState;
        bytes32 consensusState;
        bytes32 header;
    }

    protoTypes pts;

    constructor() public {
        // TODO The typeUrl should be defined in types/ethmultisig.sol
        // The schema of typeUrl follows cosmos/cosmos-sdk/codec/types/any.go
        pts = protoTypes({
            clientState: keccak256(abi.encodePacked("/ibc.lightclients.ethmultisig.v1.ClientState")),
            consensusState: keccak256(abi.encodePacked("/ibc.lightclients.ethmultisig.v1.ConsensusState")),
            header: keccak256(abi.encodePacked("/ibc.lightclients.ethmultisig.v1.Header"))
        });
    }

    function getConsensusHeight(uint64 height) internal pure returns (Height.Data memory) {
      return Height.Data({revision_number: 0, revision_height: 1});
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string memory clientId,
        uint64 height
    ) public override view returns (uint64, bool) {
        (ConsensusState.Data memory consensusState, bool found) = getConsensusState(host, clientId, getConsensusHeight(height));
        if (!found) {
            return (0, false);
        }
        return (consensusState.timestamp, true);
    }

    function getLatestHeight(
        IBCHost host,
        string memory clientId
    ) public override view returns (uint64, bool) {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        if (!found) {
            return (0, false);
        }
        return (clientState.latest_height.revision_height, true);
    }

    /**
     * @dev checkHeaderAndUpdateState checks if the provided header is valid
     */
    function checkHeaderAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory headerBytes
    ) public virtual override view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height) {
        // TODO blocked by https://github.com/hyperledger-labs/yui-ibc-solidity/issues/50
        revert("UpdateClient is not supported");
    }

    function verifySignature(ConsensusState.Data memory consensusState, MultiSignature.Data memory multisig, bytes memory signBytes) public pure returns (bool) {
        require(consensusState.addresses.length == multisig.signatures.length, "signatures length mismatch");

        for (uint i = 0; i < consensusState.addresses.length; i++) {
            require(multisig.signatures[i].length > 0, "signature is empty");
            address addr = ECRecovery.recover(keccak256(signBytes), multisig.signatures[i]);
            require(consensusState.addresses[i].toAddress() == addr, "signer mismatch");
        }

        return true;
    }

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        bytes memory clientStateBytes // serialized with pb
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        bool found;
        (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
        require(found, "consensus state not found");
        return verifySignature(
            consensusState,
            MultiSignature.decode(proof),
            makeClientStateSignBytes(height, MultiSignature.decode(proof).timestamp, consensusState.diversifier, counterpartyClientIdentifier, clientStateBytes, prefix)
        );
    }

    function verifyClientConsensusState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        string memory counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes memory prefix,
        bytes memory proof,
        bytes memory consensusStateBytes // serialized with pb
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        bool found;
        (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
        require(found, "consensus state not found");
        return verifySignature(
            consensusState,
            MultiSignature.decode(proof),
            makeConsensusStateSignBytes(height, MultiSignature.decode(proof).timestamp, consensusState.diversifier, counterpartyClientIdentifier, consensusHeight, consensusStateBytes, prefix)
        );
    }

    function verifyConnectionState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory connectionBytes // serialized with pb
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        bool found;
        (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
        require(found, "consensus state not found");
        return verifySignature(
            consensusState,
            MultiSignature.decode(proof),
            makeConnectionStateSignBytes(height, MultiSignature.decode(proof).timestamp, consensusState.diversifier, connectionId, connectionBytes, prefix)
        );
    }

    function verifyChannelState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes // serialized with pb
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        bool found;
        (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
        require(found, "consensus state not found");
        bytes memory signBytes = makeChannelStateSignBytes(height, MultiSignature.decode(proof).timestamp, consensusState.diversifier, portId, channelId, channelBytes, prefix);
        return verifySignature(
            consensusState,
            MultiSignature.decode(proof),
            signBytes
        );
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        {
            bool found;
            (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
            require(found, "consensus state not found");
        }
        bytes memory signBytes = makePacketSignBytes(height, MultiSignature.decode(proof).timestamp, consensusState.diversifier, IBCIdentifier.packetCommitmentKey(portId, channelId, sequence), commitmentBytes, prefix);
        return verifySignature(
            consensusState,
            MultiSignature.decode(proof),
            signBytes
        );
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes memory acknowledgement
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        {
            bool found;
            (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
            require(found, "consensus state not found");
        }
        bytes memory signBytes = makePacketAcknowledgementSignBytes(height, MultiSignature.decode(proof).timestamp, consensusState.diversifier, IBCIdentifier.packetAcknowledgementCommitmentKey(portId, channelId, sequence), acknowledgement, prefix);
        return verifySignature(
            consensusState,
            MultiSignature.decode(proof),
            signBytes
        );
    }

    function getClientState(IBCHost host, string memory clientId) public virtual view returns (ClientState.Data memory clientState, bool found) {
      bytes memory clientStateBytes;
      (clientStateBytes, found) = host.getClientState(clientId);
      if (!found) {
        return (clientState, false);
      }
      return (ClientState.decode(Any.decode(clientStateBytes).value), true);
    }

    function getConsensusState(IBCHost host, string memory clientId, Height.Data memory height) public virtual view returns (ConsensusState.Data memory consensusState, bool found) {
      bytes memory consensusStateBytes;
      (consensusStateBytes, found) = host.getConsensusState(clientId, height.revision_height);
      if (!found) {
        return (consensusState, false);
      }
      return (ConsensusState.decode(Any.decode(consensusStateBytes).value), true);
    }

    function parseHeader(bytes memory headerBytes) internal view returns (Header.Data memory) {
      Any.Data memory any = Any.decode(headerBytes);
      require(keccak256(abi.encodePacked(any.type_url)) == pts.header, "invalid header type");
      return Header.decode(any.value);
    }

    function makeClientStateSignBytes(
        uint64 height,
        uint64 timestamp,
        string memory diversifier,
        string memory clientID,
        bytes memory clientState,
        bytes memory prefix
    ) public pure returns (bytes memory) {
      return SignBytes.encode(
        SignBytes.Data({
          height: Height.Data({revision_number: 0, revision_height: height}),
          timestamp: timestamp,
          diversifier: diversifier,
          data_type: SignBytes.DataType.DATA_TYPE_CLIENT_STATE,
          data: StateData.encode(
            StateData.Data(
              {
                path: abi.encodePacked(prefix, IBCIdentifier.clientCommitmentKey(clientID)),
                value: clientState
              }
            )
          )
        })
      );
    }

    function makeConsensusStateSignBytes(
        uint64 height,
        uint64 timestamp,
        string memory diversifier,
        string memory clientID,
        uint64 consensusHeight,
        bytes memory consensusState,
        bytes memory prefix
    ) public pure returns (bytes memory) {
        return SignBytes.encode(
            SignBytes.Data({
                height: Height.Data({revision_number: 0, revision_height: height}),
                timestamp: timestamp,
                diversifier: diversifier,
                data_type: SignBytes.DataType.DATA_TYPE_CONSENSUS_STATE,
                data: StateData.encode(
                    StateData.Data({
                        path: abi.encodePacked(prefix, IBCIdentifier.consensusCommitmentKey(clientID, consensusHeight)),
                        value: consensusState
                    })
                )
            })
        );
    }

    function makeConnectionStateSignBytes(
        uint64 height,
        uint64 timestamp,
        string memory diversifier,
        string memory connectionID,
        bytes memory connection,
        bytes memory prefix
    ) public pure returns (bytes memory) {
        return SignBytes.encode(
            SignBytes.Data({
                height: Height.Data({revision_number: 0, revision_height: height}),
                timestamp: timestamp,
                diversifier: diversifier,
                data_type: SignBytes.DataType.DATA_TYPE_CONNECTION_STATE,
                data: StateData.encode(
                    StateData.Data({
                        path: abi.encodePacked(prefix, IBCIdentifier.connectionCommitmentKey(connectionID)),
                        value: connection
                    })
                )
            })
        );
    }

    function makeChannelStateSignBytes(
        uint64 height,
        uint64 timestamp,
        string memory diversifier,
        string memory portID,
        string memory channelID,
        bytes memory channel,
        bytes memory prefix
    ) public pure returns (bytes memory) {
        return SignBytes.encode(
            SignBytes.Data({
                height: Height.Data({revision_number: 0, revision_height: height}),
                timestamp: timestamp,
                diversifier: diversifier,
                data_type: SignBytes.DataType.DATA_TYPE_CHANNEL_STATE,
                data: StateData.encode(
                    StateData.Data({
                        path: abi.encodePacked(prefix, IBCIdentifier.channelCommitmentKey(portID, channelID)),
                        value: channel
                    })
                )
            })
        );
    }

    function makePacketSignBytes(
        uint64 height,
        uint64 timestamp,
        string memory diversifier,
        bytes32 path,
        bytes32 packetCommitment,
        bytes memory prefix
    ) public pure returns (bytes memory) {
        return SignBytes.encode(
            SignBytes.Data({
                height: Height.Data({revision_number: 0, revision_height: height}),
                timestamp: timestamp,
                diversifier: diversifier,
                data_type: SignBytes.DataType.DATA_TYPE_PACKET_COMMITMENT,
                data: StateData.encode(
                    StateData.Data({
                        path: abi.encodePacked(prefix, path),
                        value: abi.encodePacked(packetCommitment)
                    })
                )
            })
        );
    }

    function makePacketAcknowledgementSignBytes(
        uint64 height,
        uint64 timestamp,
        string memory diversifier,
        bytes32 path,
        bytes memory packetAcknowledgement,
        bytes memory prefix
    ) public pure returns (bytes memory) {
        return SignBytes.encode(
            SignBytes.Data({
                height: Height.Data({revision_number: 0, revision_height: height}),
                timestamp: timestamp,
                diversifier: diversifier,
                data_type: SignBytes.DataType.DATA_TYPE_PACKET_ACKNOWLEDGEMENT,
                data: StateData.encode(
                    StateData.Data({
                        path: abi.encodePacked(prefix, path),
                        value: abi.encodePacked(sha256(packetAcknowledgement))
                    })
                )
            })
        );
    }
}
