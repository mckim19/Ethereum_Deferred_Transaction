pragma solidity ^0.5.0;

contract TicketSeller{
  uint256 a=0;
  uint256 b=0;

  function add() public {
      lock(13);
      a+=1;
      unlock(13);
  }
  function read() public view returns(uint256) {
      return a;
  }
}

