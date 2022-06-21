const { ethers } = require('hardhat')

async function main() {
  const SuperMarioWorld = await ethers.getContractFactory("SuperMarioWorld")
  
  const superMarioWorld = await SuperMarioWorld.deploy("SuperLuigiWorld", "SPRL")
  
  await superMarioWorld.deployed()
  
  console.log("success! contract was deployed to ", superMarioWorld.address)
  
  await superMarioWorld.mint("https://ipfs.of/ipfs/QmZMK3QkG2tLiR9BkPBXrWrHQrjJGdUMWDT91zNA5Jx8QF")
  console.log('ntf successfully minted')
}


main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
