// deploy code will go here
const HDWalletProvider = require('@truffle/hdwallet-provider')
const Web3 = require('web3')
const {interface,bytecode}=require('./compile')

const provider = new HDWalletProvider('magic ribbon adapt unveil acid fence inmate earn blue distance taxi diet',
	'https://rinkeby.infura.io/v3/49834aa57ce4481abf60be06b4d934fb'	
)
	//console.log(Web3)
const web3 = new Web3(provider)

const deploy = async () => {
	//console.log(Web3)
	const accounts = await web3.eth.getAccounts()
	console.log('Attempting to deploy from', accounts[0])
	const rs= await new web3.eth.Contract(JSON.parse(interface))
		.deploy({ data: bytecode, arguments: ['Hi there'] })
		.send({ gas: '1000000', from: accounts[0] })
	console.log('Deployed to ', rs.options.address)
	provider.engine.stop()
}

deploy()
