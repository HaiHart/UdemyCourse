pragma solidity ^0.8.2;

import "./ERC721.sol";

contract  SuperMarioWorld is ERC721{
	string public name;
	string public symbol;
	uint256 public tokenCount;
	mapping(uint256=>string)private _tokenURIs;

	constructor(string memory _name, string memory _symbol){
		name=_name;
		symbol=_symbol;
	}

	modifier restrictionToTokenId(uint256 tokenId) {
		require(_owners[tokenId]!=address(0),"Token Id does not exist");
		_;
	}

	function tokenURI(uint tokenId)public view  restrictionToTokenId(tokenId) returns(string memory){
		//returns a url that points to the metadata
		return _tokenURIs[tokenId];
	}

	function mint(string memory _tokenURI)public {
		//create a new nft inside collection
		tokenCount+=1;//token Id
		_balances[msg.sender]+=1;
		_owners[tokenCount]=msg.sender;
		_tokenURIs[tokenCount]=_tokenURI;

		emit Transfer(address(0), msg.sender, tokenCount);
	}

	function supportsInterface(bytes4 intefaceId)public pure override returns(bool){
		return intefaceId== 0x80ac58cd|| intefaceId==0x5b5e139f;
	}
}