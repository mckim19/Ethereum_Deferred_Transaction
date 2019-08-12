pragma solidity ^0.5.0;

contract TicketSeller{
  uint256 a;
  uint256 b;

  constructor() public {
    a=0;
    b=0;
  }
  
  function add() public {
      lock(13);
      a+=1;
      unlock(13);
  }
  function add2() public {
	lock(13);
	a-=1;
	unlock(13);
  }
  function add3() public{
	lock(15);
	b+=2;
	unlock(15);
  }
  function read() public view returns(uint256) {
      return a;
  }
}

