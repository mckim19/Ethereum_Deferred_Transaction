pragma solidity ^0.5.0;

contract TicketSeller{
  uint256[2] a;

  constructor() public {
    a[0]=0;
    a[1]=0;
  }
  
  function add() public {
      lock(10);
      a[0]+=1;
      unlock(10);
  }
  function add2() public {
	lock(20);
	a[1]+=1;
	unlock(20);
  }
  function read() public view returns(uint256, uint256) {
      return (a[0],a[1]);
  }
}

