const assert = require('assert')
const ganache = require('ganache-cli')
const Web3 =require('web3')
const web3 = new Web3(ganache.provider())

const { interface, bytecode } = require('../compile')

let lottery
let accounts

beforeEach(async () => {
	accounts = await web3.eth.getAccounts()
	lottery = await new web3.eth.Contract(JSON.parse(interface))
		.deploy({ data: bytecode })
		.send({ from: accounts[0], gas: '1000000' })
})

describe('lottery contract', () => {
	it('Contract deploy', () => {
		assert.ok(lottery.options.address)
	})
	it('allows 1 enter', async () => {
		await lottery.methods.enter().send({
			from: accounts[0],
			value: web3.utils.toWei('0.02', 'ether')
		})
		const players = await lottery.methods.getPlayers().call({
			from: accounts[0]
		})
		assert.equal(accounts[0], players[0])
		assert.equal(1, players.length)
	})
	it('allow multiplayer', async () => {
		await lottery.methods.enter().send({
			from: accounts[0],
			value: web3.utils.toWei('0.02', 'ether')
		})
		await lottery.methods.enter().send({
			from: accounts[1],
			value: web3.utils.toWei('0.02', 'ether')
		})
		await lottery.methods.enter().send({
			from: accounts[2],
			value: web3.utils.toWei('0.02', 'ether')
		})
		const players = await lottery.methods.getPlayers().call({
			from: accounts[0]
		})
		assert.equal(accounts[0], players[0])
		assert.equal(accounts[1], players[1])
		assert.equal(accounts[2], players[2])
		assert.equal(3, players.length)
	})
	it('Require min eth', async () => {
		try {
			await lottery.methods.enter().send({
				from: accounts[0],
				value: 200
			})
			assert(false)
		} catch (err) {
			assert.ok(err)
		}
	})
	it('Not manager', async () => {
		try {
			await lottery.methods.pickWinner().send({
				from: accounts[1],
			})
			assert(false)
		} catch (err) {
			assert(err)
		}
	})
	it('end to end process', async () => {
		await lottery.methods.enter().send({
			from: accounts[0],
			value:web3.utils.toWei('2','ether')
		})
		const initialBalance = await web3.eth.getBalance(accounts[0])
		await lottery.methods.pickWinner().send({
			from:accounts[0]		
		})
		const finalBalance = await web3.eth.getBalance(accounts[0])
		const diff = finalBalance - initialBalance
		assert(diff> web3.utils.toWei('1.8','ether'))
	})
	it('empty player', async () => {
		await lottery.methods.enter().send({
			from: accounts[0],
			value: web3.utils.toWei('2', 'ether')
		})
		
		await lottery.methods.pickWinner().send({
			from: accounts[0]
		})
		const players= await lottery.methods.getPlayers().call()
		assert.equal(0,players)
	})
	it('empty balance', async () => {
		await lottery.methods.enter().send({
			from: accounts[0],
			value: web3.utils.toWei('2', 'ether')
		})

		await lottery.methods.pickWinner().send({
			from: accounts[0]
		})
		const balance = await web3.eth.getBalance(lottery.options.address)
		assert.equal(0, balance)
	})
})