const IBCHost = artifacts.require("IBCHost");
// const IBCClient = artifacts.require("IBCClient");
// const IBCConnection = artifacts.require("IBCConnection");
// const IBCChannel = artifacts.require("IBCChannel");
// const IBCHandler = artifacts.require("IBCHandler");
// const IBCMsgs = artifacts.require("IBCMsgs");
const IBCIdentifier = artifacts.require("IBCIdentifier");
const MultisigClient = artifacts.require("MultisigClient");

const deployCore = async (deployer) => {
  await deployer.deploy(IBCIdentifier);
//   await deployer.link(IBCIdentifier, [IBCHost, IBCHandler]);

//   await deployer.deploy(IBCMsgs);
//   await deployer.link(IBCMsgs, [
//     IBCClient,
//     IBCConnection,
//     IBCChannel,
//     IBCHandler
//   ]);

//   await deployer.deploy(IBCClient);
//   await deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannel]);

//   await deployer.deploy(IBCConnection);
//   await deployer.link(IBCConnection, [IBCHandler, IBCChannel]);

//   await deployer.deploy(IBCChannel);
//   await deployer.link(IBCChannel, [IBCHandler]);

  await deployer.link(IBCIdentifier, [IBCHost, MultisigClient]);
  await deployer.deploy(MultisigClient);

  await deployer.deploy(IBCHost);
};

module.exports = async function (deployer) {
  await deployCore(deployer);
};
