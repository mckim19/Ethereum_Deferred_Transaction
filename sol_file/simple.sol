pragma solidity ^0.4.24;

contract SimpleContract{
	uint256 task=0;
	uint total_executor_num = 4;
    mapping(address => uint) executor;
    string[5] book;
    constructor() public {
        book[1] = "Car Car River";
        book[2] = "River Car River";
        book[3] = "Deer River Bear";
        book[4] = "Deer Bear River";
    }
	function applys() public{
        require(total_executor_num > 0 && executor[msg.sender] == 0, "error");
        executor[msg.sender] = total_executor_num;
        total_executor_num--;
    }
	function initnode(uint64 port, uint64 peerNum, uint64 totalPeerNum) public view returns (uint64){
		require(executor[msg.sender] != 0, "error");
		init(port, peerNum, totalPeerNum);
		return port;
	}
	function senddata() public view {
		require(executor[msg.sender] != 0, "error");
		bytes memory a = bytes(book[executor[msg.sender]]);
		send(a);
	}
	function recvdata() public view returns (bytes memory){
		require(executor[msg.sender] != 0, "error");
		bytes memory b = new bytes(30);
		recv(b);
		return b;
	}
	function getTask() public view returns (string memory) {
		return book[executor[msg.sender]];
	}
}