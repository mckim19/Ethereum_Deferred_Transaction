pragma solidity ^0.5.0;

contract Viewfunction{

    function read() public view returns(uint256){
    	uint256 a = 5;
        lock(2);
        a+=2;
        unlock(2);
        return a;
    }
}
