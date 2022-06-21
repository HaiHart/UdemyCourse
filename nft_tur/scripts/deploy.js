const { ethers } = require('hardhat')

async function main() {
  const SuperMarioWorld = await ethers.getContractFactory("SuperLuigiWorldOZ")
  
  const superMarioWorld = await SuperMarioWorld.deploy("SuperLuigiWorldOZ", "SPRMO")
  
  await superMarioWorld.deployed()
  
  console.log("success! contract was deployed to ", superMarioWorld.address)
  
  await superMarioWorld.mint("https://ipfs.io/ipfs/QmNhZgsYAFsPxdEKXnvEJaurWswjjKNeBNiUM2hWevUyr8")
  console.log('ntf successfully minted')
}


main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
