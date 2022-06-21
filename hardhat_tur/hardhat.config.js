require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

const URL = "https://rinkeby.infura.io/v3/49834aa57ce4481abf60be06b4d934fb"

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: "0.8.4",
  networks: {
    rinkeby: {
      url: URL,
      accounts: ["0x72fc2d5b7bea0c0b89248beac58de52d3c86fea46e1b977aa5f4c164027be91f"]
    }
  }
};
