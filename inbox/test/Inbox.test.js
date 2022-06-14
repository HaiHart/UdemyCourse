// contract test code will go here
const assert = require('assert')
const ganache = require('ganache-cli')

const Web3 = require('web3')
const web3 = new Web3(ganache.provider())
const {interface,bytecode}=require('../compile')

// let car
// class Car{
// 	park() {
// 		return 'stopped'
// 	}
// 	drive() {
// 		return 'vroom'
// 	}
// }

// beforeEach(() => {
// 	car = new Car()
// 	console.log('a')
// })

// describe('Car', () => {
// 	it('Park function', () => {
// 		assert.equal(car.park(),'stopped')
// 	})
// 	it('Can drive', () => {
// 		assert.equal(car.drive(),'NOOOO')
// 	})
// })

let accounts
let inbox

beforeEach(async () => {
	accounts = await web3.eth.getAccounts()
	inbox=await new web3.eth.Contract(JSON.parse(interface))
		.deploy({ data: bytecode, arguments: ['Hi there'] })
		.send({ from: accounts[0], gas: '1000000' })
})

describe('Inbox', () => {
	it('deploy', () => {
		assert.ok(inbox.options.address)
	})
	it('has a default message', async () => {
		const message = await inbox.methods.message().call()
		assert.equal(message,'Hi there')
	})
	it('Can change', async () => {
		await inbox.methods.setMessage('New message').send({ from: accounts[0] })
		const message = await inbox.methods.message().call()
		assert.equal(message,'New message')
	})
})

