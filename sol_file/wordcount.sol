pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;
import "hashmap.sol";
contract wordcount {
    using strings for *;
    using hashmap for *;
    uint total_executor_num = 4;
    mapping(address => uint) executor;
    function applys() public{
        require(total_executor_num > 0 && executor[msg.sender] == 0, "error");
        executor[msg.sender] = total_executor_num;
        total_executor_num--;
    }
    function getTask() public view returns (uint){
        return executor[msg.sender];
    }
    function runnode(uint64 port, uint64 peerNum, uint64 totalPeerNum) public view returns (uint64){
        require(executor[msg.sender] != 0, "error");
        run(port, peerNum, totalPeerNum);
        return port;
    }
    function mapper(string line) public view returns (bytes memory){
        require(executor[msg.sender] != 0, "error");
        // 1. declaration
        uint size = 0;
        hashmap.BucketElement[10] memory h;
        bytes memory bts;
        if (strings.equals(line.toSlice(), "".toSlice())){
            write(bts);
            return bts;
        }
        // 2. string tokenizer
        strings.slice memory s = line.toSlice();
        strings.slice memory delim = " ".toSlice();
        string[] memory parts = new string[](s.count(delim)+1);
        uint i = 0;
        for(i = 0; i < parts.length; i++){
            parts[i] = s.split(delim).toString();
        }
        // 3. count the number of words
        uint32 val = 1;
        uint j = 0;
        for(i = 0; i<parts.length; i++){
            //insert element in the hashmap
            uint idx = (bytes(parts[i]).length)%10;
            if(h[idx].val == 0){
                size++;
                hashmap.Node[] memory arr1 = new hashmap.Node[](20);
                h[idx] = hashmap.BucketElement(val, parts[i], arr1, 0);
            }
            else if(strings.equals(h[idx].word.toSlice(), parts[i].toSlice())){
                h[idx].val += val;
            }
            else{
                for(j = 0; j<h[idx].cnt; j++){
                     if(strings.equals(h[idx].list[j].word.toSlice(), parts[i].toSlice()))
                         break;
                }
                if(j==h[idx].cnt){  // not found
                    size++;
                    h[idx].cnt++;
                    h[idx].list[j] = hashmap.Node(1, parts[i]);
                } else {  //found
                    h[idx].list[j].val += val;
                }
            }
        }
        // 4. convert hashmap to Node[], then convert Node[] to bytes
        bts = hashmap.nodeArrayToBytes(hashmap.hashmapToNodeArray(h, size));
        write(bts);
        return bts;
        // return h;
    }
    function shfflerAndReducer(uint mappernum) public view returns (hashmap.BucketElement[10] memory){
        require(executor[msg.sender]!=0, "error");

        hashmap.BucketElement[10] memory h;
        uint size = 0;
        bytes memory data = new bytes(100);
        hashmap.Node[] memory ns;
        uint idx = 0;
        uint i = 0;
        uint j = 0;

        uint receivedMapperNum = 0;
        for(receivedMapperNum = 0; receivedMapperNum<mappernum;){
            read(data);
            if (data.length==0) {
                receivedMapperNum++;
            }
            if(receivedMapperNum==mappernum){
                break;
            }
            ns = hashmap.bytesToNodeArray(data);
            // insert element in the hashmap
            for (i = 0; i<ns.length; i++){
                idx = (bytes(ns[i].word).length)%10;
                if(h[idx].val == 0){
                    size++;
                    hashmap.Node[] memory arr1 = new hashmap.Node[](20);
                    h[idx] = hashmap.BucketElement(ns[i].val, ns[i].word, arr1, 0);
                }
                else if(strings.equals(h[idx].word.toSlice(), ns[i].word.toSlice())){
                    h[idx].val += ns[i].val;
                }
                else{
                    for(j = 0; j<h[idx].cnt; j++){
                         if(strings.equals(h[idx].list[j].word.toSlice(), ns[i].word.toSlice()))
                             break;
                    }
                    if(j==h[idx].cnt){  // not found
                        size++;
                        h[idx].cnt++;
                        h[idx].list[j] = hashmap.Node(1, ns[i].word);
                    }
                    else{  //found
                        h[idx].list[j].val += ns[i].val;
                    }
                }
            }
        }
        return h;
    }

    function submit() public view {
        require(executor[msg.sender] != 0, "error");
    }
}