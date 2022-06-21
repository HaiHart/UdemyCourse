const fs=require('fs')

async function main() {
	const [developer] = await ethers.getSigners()
	console.log(`Deploying contracts with accounts: ${developer.address}`)
	
	const balance = await developer.getBalance()
	console.log(`Account Balance: ${balance.toString()}`)
	
	const Token = await ethers.getContractFactory('Token')
	const token = await Token.deploy()
	console.log(`Token address: ${token.address}`)
	
	const data = {
		address: token.address,
		abi:JSON.parse(token.interface.format('json'))
	}
	fs.writeFileSync('frontend/src/Token.json',JSON.stringify(data))
}

main().then(() => { process.exit(0) }).catch((err) => { console.log(err); process.exit(1); })