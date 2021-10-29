package types

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

func (cs *ClientState) ClientType() string {
	return ClientType
}

func (cs *ClientState) GetLatestHeight() exported.Height {
	return clienttypes.Height(cs.LatestHeight)
}

func (cs *ClientState) Validate() error {
	return nil
}

func (cs *ClientState) GetProofSpecs() []*ics23.ProofSpec {
	return nil
}

// Initialization function
// Clients must validate the initial consensus state, and may store any client-specific metadata
// necessary for correct light client operation
func (cs *ClientState) Initialize(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.ConsensusState) error {
	return nil
}

// Status function
// Clients must return their status. Only Active clients are allowed to process packets.
func (cs *ClientState) Status(ctx sdk.Context, clientStore sdk.KVStore, cdc codec.BinaryCodec) exported.Status {
	if !clienttypes.Height(cs.FrozenHeight).IsZero() {
		return exported.Frozen
	}

	return exported.Active
}

// Genesis function
func (cs *ClientState) ExportMetadata(_ sdk.KVStore) []exported.GenesisMetadata {
	panic("not implemented") // TODO: Implement
}

// Update and Misbehaviour functions
func (cs *ClientState) CheckHeaderAndUpdateState(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Header) (exported.ClientState, exported.ConsensusState, error) {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) CheckMisbehaviourAndUpdateState(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Misbehaviour) (exported.ClientState, error) {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) CheckSubstituteAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore sdk.KVStore, substituteClientStore sdk.KVStore, substituteClient exported.ClientState) (exported.ClientState, error) {
	panic("not implemented") // TODO: Implement
}

// Upgrade functions
// NOTE: proof heights are not included as upgrade to a new revision is expected to pass only on the last
// height committed by the current revision. Clients are responsible for ensuring that the planned last
// height of the current revision is somehow encoded in the proof verification process.
// This is to ensure that no premature upgrades occur, since upgrade plans committed to by the counterparty
// may be cancelled or modified before the last planned height.
func (cs *ClientState) VerifyUpgradeAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, store sdk.KVStore, newClient exported.ClientState, newConsState exported.ConsensusState, proofUpgradeClient []byte, proofUpgradeConsState []byte) (exported.ClientState, exported.ConsensusState, error) {
	panic("not implemented") // TODO: Implement
}

// Utility function that zeroes out any client customizable fields in client state
// Ledger enforced fields are maintained while all custom fields are zero values
// Used to verify upgrades
func (cs *ClientState) ZeroCustomFields() exported.ClientState {
	return &ClientState{
		LatestHeight: cs.LatestHeight,
	}
}

// State verification functions
func (cs ClientState) VerifyClientState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, counterpartyClientIdentifier string, proof []byte, clientState exported.ClientState) error {
	cons, sigData, err := produceVerificationArgs(store, cdc, cs, height, prefix, proof)
	if err != nil {
		return err
	}
	path, err := ClientCommitmentKey(prefix.Bytes(), counterpartyClientIdentifier)
	if err != nil {
		return err
	}
	signBz, err := ClientStateSignBytes(cdc, height.(clienttypes.Height), sigData.Timestamp, cons.Diversifier, path, clientState)
	if err != nil {
		return err
	}
	return VerifySignature(cons.GetAddresses(), sigData, signBz)
}

func (cs ClientState) VerifyClientConsensusState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, counterpartyClientIdentifier string, consensusHeight exported.Height, prefix exported.Prefix, proof []byte, consensusState exported.ConsensusState) error {
	cons, sigData, err := produceVerificationArgs(store, cdc, cs, height, prefix, proof)
	if err != nil {
		return err
	}
	path, err := ConsensusCommitmentKey(prefix.Bytes(), counterpartyClientIdentifier, consensusHeight.GetRevisionHeight())
	if err != nil {
		return err
	}
	signBz, err := ConsensusStateSignBytes(cdc, height.(clienttypes.Height), sigData.Timestamp, cons.Diversifier, path, consensusState)
	if err != nil {
		return err
	}
	return VerifySignature(cons.GetAddresses(), sigData, signBz)
}

func (cs *ClientState) VerifyConnectionState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, connectionID string, connectionEnd exported.ConnectionI) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyChannelState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, portID string, channelID string, channel exported.ChannelI) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyPacketCommitment(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64, commitmentBytes []byte) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyPacketAcknowledgement(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64, acknowledgement []byte) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyPacketReceiptAbsence(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyNextSequenceRecv(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, nextSequenceRecv uint64) error {
	panic("not implemented") // TODO: Implement
}

// produceVerificationArgs perfoms the basic checks on the arguments that are
// shared between the verification functions and returns the public key of the
// consensus state, the unmarshalled proof representing the signature and timestamp
// along with the solo-machine sequence encoded in the proofHeight.
func produceVerificationArgs(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	cs ClientState,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
) (*ConsensusState, *MultiSignature, error) {

	var multiSig MultiSignature
	if err := cdc.Unmarshal(proof, &multiSig); err != nil {
		return nil, nil, err
	}

	cons, err := getConsensusState(store, cdc, height)
	if err != nil {
		return nil, nil, err
	}

	if cons.GetTimestamp() > multiSig.Timestamp {
		return nil, nil, sdkerrors.Wrapf(ErrInvalidProof, "the consensus state timestamp is greater than the signature timestamp (%d >= %d)", cons.GetTimestamp(), multiSig.Timestamp)
	}

	return cons, &multiSig, nil
}
