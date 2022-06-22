const { ethers } = require('hardhat')

async function main() {
  const SuperMarioWorld = await ethers.getContractFactory("SuperMarioWorldERC1155")
  
  const superMarioWorld = await SuperMarioWorld.deploy("SuperMarioWorldERC1155", "SPRM1155")
  
  await superMarioWorld.deployed()
  
  console.log("success! contract was deployed to ", superMarioWorld.address)
  
  await superMarioWorld.mint(15,"https://ipfs.io/ipfs/QmUXfU6VBP8gjX6ydZSZLpNeqoNtkgrxhKdHPVBtPD8Nne")
  console.log('ntf successfully minted')
}


main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
