const { expect } = require('chai')

describe('Token contract', () => {
	let owner, addr1, addr2,Token,token;

	beforeEach(async () => {
		
		Token = await ethers.getContractFactory('Token')
		token = await Token.deploy()
		const [addr0, addr_1, addr_2, _] = await ethers.getSigners();
		owner = addr0
		addr1 = addr_1
		addr2=addr_2
	})
	describe('Deployment', () => {
		it('should set right owner', async()=> {
			expect(await token.owner()).to.equal(owner.address)
		})
		it('should assign the total supplt of token to owner', async () => {
			const ownerBalance = await token.balanceOf(owner.address)
			expect(await token.totalSuppl()).to.equal(ownerBalance)
		})
	})
	describe('Transaction', () => {
		it('Should transfer token between account', async () => {
			await token.transfer(addr1.address, 50)
			const addrBalance = await token.balanceOf(addr1.address)
			expect(addrBalance).to.equal(50);
			
			await token.connect(addr1).transfer(addr2.address, 50)
			const addr2Balance = await token.balanceOf(addr2.address)
			expect(addr2Balance).to.equal(50)
			
		})
		it('Should for lack of token ', async () => {
			const initialBalanceOwner = await token.balanceOf(owner.address)
			
			await expect(
				token.connect(addr1).transfer(owner.address,1)
			).to.be.revertedWith('Not enough token')
			expect(
				await token.balanceOf(owner.address)
			).to.equal(initialBalanceOwner)
			
		})
		it('Should update balance after transfer', async () => {
			const initialBalanceOwner = await token.balanceOf(owner.address)
			
			await token.transfer(addr1.address, 100)
			await token.transfer(addr2.address, 50)
			const finalBalance = await token.balanceOf(owner.address)
			expect(finalBalance).to.equal(initialBalanceOwner - 150)
			
			const addr1Balance = await token.balanceOf(addr1.address)
			expect(addr1Balance).to.equal(100)
			
			const addr2Balance = await token.balanceOf(addr2.address)
			expect(addr2Balance).to.equal(50)
		})
	})
})

