pragma solidity ^0.4.17;



contract Campaign{
    address public manager;
    uint public minimumContribution;
    mapping(address=>bool) public contributors;
    uint public contributorCount;
    struct Request{
        string description;
        uint value;
        address recipient;
        bool complete;
        mapping(address=>bool) approvals;
        uint apporvalsCount;
    }
    Request[] public requests;

    modifier restrict(){
        require(msg.sender==manager);
        _;
    }

    function Campaign(uint mini,address creator)public{
        manager=creator;
        minimumContribution=mini;
        contributorCount=0;
    }

    function contribute()public payable{
        require(msg.value>minimumContribution);
        contributors[msg.sender]=true;
        contributorCount++;
    }

    function createRequest(string des,uint value, address rep)public restrict{
        Request memory newRequest= Request({
            description:des,
            value:value,
            recipient:rep,
            complete:false,
            apporvalsCount:0
        });
        requests.push(newRequest);
    }

    function approveRequest(uint index)public{
        Request storage req= requests[index];

        require(contributors[msg.sender]);
        require(!requests[index].approvals[msg.sender]);

        req.approvals[msg.sender]=true;
        req.apporvalsCount++;
    }

    function finalizeRequest(uint index)public restrict {
        Request storage req=requests[index];
        require(!req.complete);
        require(req.apporvalsCount>(contributorCount/2));
        req.recipient.transfer(req.value);
        req.complete=true;

    }

}

contract Factory{
    address[] public deployedCampaign;

    function createCampaign(uint mini)public{
        address newCampaign = new Campaign(mini,msg.sender); 
        deployedCampaign.push(newCampaign);
    }

    function getDeployed()public view returns(address[]){
        return deployedCampaign;
    }

}