pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;
import "strings.sol";

library hashmap {
    using strings for *;

    struct BucketElement{
        uint32 val;
        string word;
        Node[] list;
        uint32 cnt;
    }
    struct Node{
        uint32 val;
        string word;
    }
    function hashmapToNodeArray(BucketElement[10] memory b, uint size) internal pure returns (Node[] memory) {
        Node[] memory ns = new Node[](size);
        uint i = 0;
        uint j = 0;
        uint idx = 0;
        for (i = 0; i<b.length; i++){
            if (b[i].val!=0) {
                ns[idx] = Node(b[i].val, b[i].word);
                idx++;
            }
            for(j = 0; j<b[i].cnt; j++){
                ns[idx] = Node(b[i].list[j].val,b[i].list[j].word);
                idx++;
            }
        }
        return ns;
    }
    function nodeArrayToBytes(Node[] memory n) internal pure returns (bytes memory) {
    // function nodeArrayToBytes(Node[] memory n) internal pure returns (uint length) {
        uint i = 0;
        uint _size = 4;
        for (i = 0; i<n.length; i++) {
            // _size = "sizeof" word + sizeof" n.val + "sizeof" n.word
            _size = _size + 1 + 4 + bytes(n[i].word).length;
        }
        bytes memory _data = new bytes(_size);
        uint counter = 0;
        uint idx = 0;
        // return bytes(n[2].word).length;
        for(i = 0; i<4; i++){
            _data[counter] = byte(uint32(n.length)>>(8*i)&uint32(255));
            counter++;
        }
        for (idx = 0; idx<n.length; idx++){
            _data[counter] = byte(int8(bytes(n[idx].word).length));
            counter++;
            for(i = 0; i<4; i++){
                _data[counter] = byte(n[idx].val>>(8*i)&uint32(255));
                counter++;
            }
            bytes memory tmpword = bytes(n[idx].word);
            for(i = 0; i<tmpword.length;i++){
                _data[counter] = tmpword[i];
                counter++;
            }
        }
        return (_data);
    }
    function bytesToNodeArray(bytes memory data) internal pure returns (Node[] memory) {
        uint i = 0;
        Node[] memory ns;
        uint32 ns_size = 0;
        uint idx = 0;

        uint32 temp;
        uint8 n_size = 0;
        uint j = 0;
        for(i = 0; i<4; i++){
            temp = uint32(data[i]);
            temp <<= 8*i;
            ns_size ^= temp;
        }
        ns = new Node[](ns_size);
        for (idx = 0; idx<ns_size; idx++) {
            n_size = uint8(data[i]);
            i++;
            for (j = 0; j<4; j++){
                temp = uint32(data[i]);
                temp <<= 8*j;
                ns[idx].val ^= temp;
                i++;
            }
            bytes memory str = new bytes(data.length);
            for (j = 0; j<n_size; j++){
                str[j] = data[i];
                i++;
            }
            ns[idx].word = string(str);
        }
        return ns;
    }
}