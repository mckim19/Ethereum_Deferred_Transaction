pragma solidity ^0.5.2;
contract example {
	uint64 num;
    function vote(uint candidate_num) public {
		require(candidate_num>3);
		num++;
		lock(num);
		unlock(num);
    }
}
