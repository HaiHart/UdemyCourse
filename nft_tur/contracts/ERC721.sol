pragma solidity ^0.8.2;

contract ERC721 {
	mapping(address=>uint256)internal _balances;
	mapping(uint256=>address)internal _owners;
	mapping(address=>mapping(address=>bool))private _operatorApprovals;
	mapping(uint256=>address) internal token_approvals;

	 event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);

    
    event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId);

    
    event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);

   
    function balanceOf(address _owner) external view returns (uint256){
		//return num of nft assign to owner
		require( _owner!=address(0) , "Address is zero");
		return _balances[_owner];
	}

    
    function ownerOf(uint256 _tokenId) external view returns (address){
		//finds the owner of the nft
		address owner=_owners[_tokenId];
		require(owner!=address(0), "Token Id does not exist");
		return owner;
	}

    function setApprovalForAll(address _operator, bool _approved) external{
		//enable or disable an operator to manage all of msg.senders asset
		_operatorApprovals[msg.sender][_operator]=_approved;
		emit ApprovalForAll(msg.sender, _operator, _approved);
	}

    function isApprovedForAll(address _owner, address _operator) public view returns (bool){
		//if and address is an operator for another address
		return _operatorApprovals[_owner][_operator];
	}

    function approve(address _approved, uint256 _tokenId) public payable{
		//update an approved address  for an nft
		address owner = _owners[_tokenId];
		require(msg.sender == owner || _operatorApprovals[owner][msg.sender], "Msg.sender is not the owner or the apporved operators");
		token_approvals[_tokenId]=_approved;
		emit Approval(owner,_approved, _tokenId);
	}

    function getApproved(uint256 _tokenId) public view returns (address){
		require(_owners[_tokenId]!=address(0),"Token id does not exist");
		return token_approvals[_tokenId];
	}

	function _checkOnERC721Received() private pure returns(bool)	{
		return true;
	}

    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes memory data) public payable{
		//standard transfer
		//check if onERC721Received is implemented WHEN sending  smart contracts
		transferFrom(_from, _to, _tokenId);
		require(_checkOnERC721Received(),"Receiver not implemented");

	}

	
   
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable{
		safeTransferFrom(_from, _to, _tokenId, "");
	}

    
    function transferFrom(address _from, address _to, uint256 _tokenId) public payable{
		//transfer ownership of an nft
		address owner=_owners[_tokenId];

		require(
			msg.sender ==owner||
			getApproved(_tokenId)==msg.sender||
			isApprovedForAll(owner,msg.sender),
			"Not allowed"
		);
		require(
			owner==_from,"From is not the owner"
		);				
		require(
			_owners[_tokenId]!=address(0),"Tokenid does not exist"
		);		
		//
		approve(address(0), _tokenId);
		//
		_balances[_from]-=1;
		_balances[_to]+=1;
		_owners[_tokenId]=_to;
		emit Transfer(_from, _to, _tokenId);
	}

   
	function supportsInterface(bytes4 interfaceID)virtual external pure returns (bool){
		//EIP 165: if a contract implements another interface
		return interfaceID==0x80ac58cd;
	}
}