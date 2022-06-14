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
		console.log(inbox)
	})

})

