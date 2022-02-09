# ibc-ethmultisig-client

This client is a dialect of https://github.com/datachainlab/ibc-multisig-client/blob/main/docs/ibc-multisig-client.md.

The main changes are as follows:

- Supported only secp256k1 key
- Modified `SignBytes` format to reduce the protobuf serialization

## Directories

- [contracts/core/ibc/lightclients](./contracts/core/ibc/lightclients): An implementation of the client in solidity
- [modules/light-clients/xx-ethmultisig](./modules/light-clients/xx-ethmultisig/): (WIP) An implementation of the client in golang
- [modules/relay/ethmultisig](./modules/relay/ethmultisig/): A relay module for the client
