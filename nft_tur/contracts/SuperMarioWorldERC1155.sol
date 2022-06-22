pragma solidity ^0.8.2;	

// import "./ERC1155.sol";

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract SuperMarioWorldERC1155 is ERC1155 {
	
	string public name;
	string public symbol;
	uint256 public tokenCount;

	mapping(uint256=>string) private tokenURIs;
	
	constructor(string memory _name, string memory _symbol){
		name=_name;
		symbol=_symbol;
	}
	function uri(uint256 tokenId)public view returns(string memory) {
		return tokenURIs[tokenId];
	}
	function mint(uint256 amount, string memory _uri)public ZeroAddress(msg.sender){
		tokenCount+=1;
		_balances[tokenCount][msg.sender]+=amount;
		tokenURIs[tokenCount]=_uri;
		emit TransferSingle(msg.sender, address(0),msg.sender, tokenCount, amount);
	}

	function supportsInterface(bytes4 _interfaceId)public pure override returns(bool) {
		return _interfaceId==0xd9b67a26||_interfaceId==0x0e89341c;
	}
}