pragma solidity ^0.5.0;

contract TicketSeller{
  uint256 a=5;
  
  function add() public
  {
	require(a>3);
	a+=1;
  }
  
  function read() public view returns(uint256) {
      	return a;
  }
}

