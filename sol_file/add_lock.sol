pragma solidity ^0.5.0;

contract TicketSeller{
  uint256 a;

  constructor() public {
    a=0;
  }
  
  function add() public {
      lock(0);
      a+=1;
      unlock(0);
  }
  
  function read() public view returns(uint256) {
      return a;
  }
}

