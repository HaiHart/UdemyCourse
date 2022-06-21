pragma solidity ^0.8.2;

contract ERC1155 {

	mapping(uint256=>mapping(address=>uint256)) internal _balances;
	mapping(address=>mapping(address=>bool))internal _operatorApprovals;

	event ApprovalForAll(address indexed owner,address indexed operator,bool approved );

	event TransferSingle(address indexed _operator,address indexed _from,address indexed _to,uint256 _id, uint256 amount);

	event TransferBatch(address _operator,address _from,address _to,uint256[]_ids,uint256[] values);

	modifier ZeroAddress(address account){
		require(account!=address(0),"Address is zero");
		_;
	}

	function balanceOf(address _account,uint256 _id) public view ZeroAddress(_account) returns(uint256){
		//getbalance of account token
		return _balances[_id][_account];
	}

	function balanceOfBatch(address[]memory _accounts,uint256[] memory _ids) public view returns(uint256[] memory) {
		//get balance of multiple account's token
		require(_accounts.length!=_ids.length,"Acc and id not equal in length");
		uint256[] memory batchBalances=new uint256[](_accounts.length);
		for (uint256 i =0; i < _accounts.length;i++) {
			batchBalances[i]=balanceOf(_accounts[i], _ids[i]);
		}
	}

	function isApprovedForAll(address _owner,address _operator) public view returns(bool) {
		// Checks if an address is an operator for another address
		return _operatorApprovals[_owner][_operator];
	}

	function setApprovalForAll(address _operator, bool _approved)public{
		// enable or disable an operator to manage all of msg.sender assets
		_operatorApprovals[msg.sender][_operator]=_approved;
		emit ApprovalForAll(msg.sender,_operator,_approved);
	}

	function _transfer(address _from, address _to,uint256 _id,uint256 amount)private {
		uint256 fromBalance=_balances[_id][_from];
		require(fromBalance>=amount,"Insufficient balance");
		_balances[_id][_from]=fromBalance-amount;
		_balances[_id][_to]+=amount;
	}

	function safeTransferFrom(address _from,address _to,uint256 _id, uint amount)public ZeroAddress(_to) virtual{
		require(_from==msg.sender||isApprovedForAll(_from,msg.sender),"Msg.sender is not the owner or approved by the owner");
		_transfer(_from, _to, _id, amount);
		emit TransferSingle(msg.sender,_from,_to,_id,amount);
		require(_checkOnERC1155Received(),"Reciver is not implemented");
	}

	function _checkOnERC1155Received()private returns(bool){
		// temp

		return true;
	}

	function safeBatchTransferFrom(address _from,address _to,uint256[] memory _ids,uint256[] memory amounts)public ZeroAddress(_to) {
		require(_from==msg.sender||isApprovedForAll(_from,msg.sender),"Msg.sender is not the owner or approved by the owner");
		require(_ids.length==amounts.length,"Ids and amounts are not the same");
		for(uint256 i=0; i<_ids.length;i++){
			uint256 _id=_ids[i];
			uint256 amount=amounts[i];
			_transfer(_from, _to, _id, amount);
		}
		emit TransferBatch(msg.sender,_from,_to,_ids,amounts);
		require(_checkOnBatchERC1155Received(),"Reciver is not implemented");
	}
	function _checkOnBatchERC1155Received()private returns(bool){
		return true;
		//temp
	}

	function supportsInterface(bytes4 interfaceId)public pure virtual returns(bool){
		return interfaceId==0xd9b67a26;
	}
}