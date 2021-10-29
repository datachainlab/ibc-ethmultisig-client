package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

var commitmentSlot = big.NewInt(0)

const (
	clientPrefix         = uint8(0)
	consensusStatePrefix = uint8(1)
	connectionPrefix     = uint8(2)
	channelPrefix        = uint8(3)
	packetPrefix         = uint8(4)
	packetAckPrefix      = uint8(5)
)

// Commitment key generator

func ClientCommitmentKey(prefix []byte, clientId string) ([]byte, error) {
	key, err := keccak256AbiEncodePacked(clientPrefix, clientId)
	if err != nil {
		return nil, err
	}
	return append(prefix, key...), nil
}

func ConsensusCommitmentKey(prefix []byte, clientId string, height uint64) ([]byte, error) {
	key, err := keccak256AbiEncodePacked(consensusStatePrefix, clientId, "/", height)
	if err != nil {
		return nil, err
	}
	return append(prefix, key...), nil
}

func ConnectionCommitmentKey(prefix []byte, connectionId string) ([]byte, error) {
	key, err := keccak256AbiEncodePacked(connectionPrefix, connectionId)
	if err != nil {
		return nil, err
	}
	return append(prefix, key...), nil
}

func ChannelCommitmentKey(prefix []byte, portId, channelId string) ([]byte, error) {
	key, err := keccak256AbiEncodePacked(channelPrefix, portId, "/", channelId)
	if err != nil {
		return nil, err
	}
	return append(prefix, key...), nil
}

func PacketCommitmentKey(prefix []byte, portId, channelId string, sequence uint64) ([]byte, error) {
	key, err := keccak256AbiEncodePacked(packetPrefix, portId, "/", channelId, "/", sequence)
	if err != nil {
		return nil, err
	}
	return append(prefix, key...), nil
}

func PacketAcknowledgementCommitmentKey(prefix []byte, portId, channelId string, sequence uint64) ([]byte, error) {
	key, err := keccak256AbiEncodePacked(packetAckPrefix, portId, "/", channelId, "/", sequence)
	if err != nil {
		return nil, err
	}
	return append(prefix, key...), nil
}

// keccak256AbiEncodePacked only covers some data types.
func keccak256AbiEncodePacked(data ...interface{}) ([]byte, error) {
	// abi.encodePacked
	var bzs [][]byte
	for _, v := range data {
		var bz []byte
		switch vt := v.(type) {
		case *big.Int:
			bz = math.U256Bytes(vt)
		case uint64, uint8:
			b := new(bytes.Buffer)
			if err := binary.Write(b, binary.BigEndian, vt); err != nil {
				return nil, err
			}
			bz = b.Bytes()
		case string:
			// XXX address may be represented by common.Address
			// also, no slot uses an address for now
			if common.IsHexAddress(vt) {
				vt = strings.TrimPrefix(vt, "0x")
				if vt == "" || vt == "0" {
					bz = []byte{0}
					break
				}
				vt = evenLengthHex(vt)
				var err error
				bz, err = hex.DecodeString(vt)
				if err != nil {
					return nil, err
				}
			} else {
				bz = []byte(vt)
			}
		case common.Address:
			bz = vt.Bytes()[:]
		case []byte:
			bz = common.RightPadBytes(vt, len(vt))
		default:
			return nil, fmt.Errorf("unsupported type for abiEncodePacked: %s", reflect.TypeOf(v))
		}
		bzs = append(bzs, bz)
	}

	// keccak256
	return crypto.Keccak256(bzs...), nil
}

func evenLengthHex(v string) string {
	if len(v)%2 == 1 {
		v = "0" + v
	}
	return v
}
