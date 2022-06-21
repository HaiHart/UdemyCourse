//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Token{
    string public name = 'My hardhat token';
    string public symbol='MHT';

    uint public totalSuppl=1000000;
    address public owner;
    mapping (address=>uint) balances;
    constructor(){
        balances [msg.sender]=totalSuppl;
        owner=msg.sender;
    }

    function transfer(address to, uint amount)  external{
        console.log('sender balance is %s tokens',balances[msg.sender]);
        console.log('Trying to send %s tokens to %s',amount,to);
        require(balances[msg.sender]>=amount,'Not enough token');
        balances[msg.sender]-=amount;
        balances[to]+=amount;
    }

    function balanceOf(address account)external view returns(uint) {
        return balances[account];
        
    }
}

