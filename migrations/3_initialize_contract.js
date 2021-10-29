const IBCHost = artifacts.require("IBCHost");

module.exports = async function (deployer, network, accounts) {
  const ibcHost = await IBCHost.deployed();
  for(const promise of [
    () => ibcHost.setIBCModule(accounts[0]),
  ]) {
    const result = await promise();
    console.log(result);
    if(!result.receipt.status) {
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
};
