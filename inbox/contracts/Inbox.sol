pragma solidity ^0.4.17;
// linter warnings (red underline) about pragma version can igonored!
   contract Inbox{
        string public message;
        function Inbox(string InitialMessage) public{
            message=InitialMessage;
        }
        function setMessage(string newMessage)public{
            message = newMessage;
        }
        function doMath( int a, int b){
            a+b;
            a-b;
            a*b;
            a==0;
        }
    }
// contract code will go here
