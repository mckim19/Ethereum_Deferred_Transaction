pragma solidity ^0.5.0;

contract TicketSeller{
	uint256 task=0;
	function add() public view returns(uint256){
		require(task<10);
		uint256 a;
		lock(13);
		a++;
		unlock(13);
	}
}
