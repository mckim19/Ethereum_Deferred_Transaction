pragma solidity ^0.5.0;

contract TicketSeller{
  uint256 a;

  constructor() public {
    a=0;
  }
  
  function add() public {
      a+=1;
  }
  
  function read() public view returns(uint256) {
      return a;
  }
}

