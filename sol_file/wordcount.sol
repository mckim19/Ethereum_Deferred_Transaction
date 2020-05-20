pragma experimental ABIEncoderV2;
pragma solidity ^0.4.24;
import "github.com/Arachnid/solidity-stringutils/strings.sol";
import "./hashmap.sol";
contract wordcount{
    using strings for *;
    using hashmap for *;
    

    uint total_executor_num=4;
    mapping(address => uint) executor;
    
    string[5] book;
    constructor() public{
        book[1] = "Car Car River";
        book[2] = "River Car River";
        book[3] = "Deer River Bear";
        book[4] = "Deer Bear River";
    }
    function applys() public{
        require(total_executor_num >0 && executor[msg.sender] == 0);
        executor[msg.sender] = --total_executor_num;
    }
    
    function execute() public view returns (hashmap.BucketElement[10] memory){
        require(executor[msg.sender]!=0);
        
        // 1. declaration
        hashmap.BucketElement[10] memory h;
        uint task = executor[msg.sender];
        
        // 2. string tokenizer
        strings.slice memory s = book[task].toSlice();
        strings.slice memory delim = " ".toSlice();
        string[] memory parts = new string[](s.count(delim)+1);
        uint i=0;
        uint j=0;
        for(i=0; i<parts.length; i++){
            parts[i] = s.split(delim).toString();
        }
        
        // 3. count the number of words

        for(i=0; i<parts.length; i++){
            uint idx = (bytes(parts[i]).length)%parts.length;
            if(h[idx].val == 0){
                hashmap.Node[] memory arr1 = new hashmap.Node[](parts.length);
                h[idx] = hashmap.BucketElement(1, parts[i], arr1, 0);
            }
            else if(strings.equals(h[idx].word.toSlice(), parts[i].toSlice())){
                h[idx].val++;
            }
            else{
                for(j=0; j<h[idx].cnt; j++){
                     if(strings.equals(h[idx].list[j].word.toSlice(), parts[i].toSlice()))
                         break;
                }
                if(j==h[idx].cnt){  // not found
                    h[idx].cnt++;
                    h[idx].list[j] = hashmap.Node(1, parts[i]);
                }
                else{  //found
                    h[idx].list[j].val++;
                }
            }
        }
        return h;
    }
    function test() public view returns (hashmap.Node){
        hashmap.BucketElement[10] memory h;
        hashmap.Node[] memory arr1 = new hashmap.Node[](5);
        h[0] = hashmap.BucketElement(1, "hello", arr1, 0);
        h[0].list[0] = hashmap.Node(1, book[1]);
        h[0].cnt++;
        return h[0].list[0];
    }
}



