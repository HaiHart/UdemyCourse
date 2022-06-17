const assert = require('assert')
const ganache = require('ganache-cli')
const Web3 = require('web3')
const web3 = new Web3(ganache.provider())

const compiledFactory = require('../ethereum/build/Factory.json')
const compiledCampaign = require('../ethereum/build/Campaign.json')

let accounts
let factory
let campaignAddress
let campaign

beforeEach(async () => {
	accounts = await web3.eth.getAccounts()
	factory = await new web3.eth.Contract(JSON.parse(compiledFactory.interface))
		.deploy({
			data: compiledFactory.bytecode
		})
		.send({
			from: accounts[1],
			gas:'1000000'
		})
	await factory.methods.createCampaign('100').send({
		from: accounts[0],
		gas: '1000000'
	})
	const addresses = await factory.methods.getDeployed().call();
	campaignAddress=addresses[0]
	campaign = await new web3.eth.Contract(JSON.parse(compiledCampaign.interface), campaignAddress)
})

describe('Campaigns', () => {
	it('Deployed factory and campaign', () => {
		assert.ok(factory.options.address)
		assert.ok(campaign.options.address)
	})
	it('marks caller is campaign manager', async () => {
		const manager = await campaign.methods.manager().call()
		assert.equal(accounts[0],manager)
	})
	it('Contributors checking', async () => {
		await campaign.methods.contribute().send({
			from: accounts[2],
			value:'101'
		})
		const isContributors = await campaign.methods.contributors(accounts[2]).call()
		assert(isContributors)
	})
	it('requires minimum contribution', async () => {
		try {
			await campaign.methods.contribute().send({
				from: accounts[3],
				value:'5'
			})
			assert(false)
		} catch (err) {
			assert(err)
		}
	})
	it('allows request', async () => {
		await campaign.methods.createRequest(
			'Buy stuff',
			'100',
			accounts[0]
		).send({
			from: accounts[0],
			gas:'1000000'
		})
		const requests = await campaign.methods.requests(0).call()
		assert.equal('Buy stuff',requests.description)
	})
	it('process request', async () => {
		await campaign.methods.contribute().send({
			from: accounts[1],
			value:web3.utils.toWei('10','ether')
		})
		await campaign.methods.createRequest(
			'A', web3.utils.toWei('5', 'ether'),accounts[1]
		)
			.send({
				from: accounts[0],
				gas:'1000000'
			})
		await campaign.methods.approveRequest(0).send({
			from: accounts[1],
			gas:'1000000'
		})
		
		await campaign.methods.finalizeRequest(0).send({
			from: accounts[0],
			gas:'1000000'
		})
		let balance = await web3.eth.getBalance(accounts[1])
		balance = web3.utils.fromWei(balance, 'ether')
		balance = parseFloat(balance)
		// console.log(balance)
		assert(balance>94)
	})
})