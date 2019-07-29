pragma solidity ^0.5.4;
//We believe voters are innocent.
contract C_v3 {
    uint constant POPULATION_NUM=100;
    uint[3] candidate;
    uint v;
    bool final_flag;
    uint m;
    function vote(uint candidate_num) public
    {
        /* x,y,z does not have to store sequence! */
        /* only v have to store sequence! */
        v++;
        if(v == POPULATION_NUM)
            final_flag = true;
        candidate[candidate_num]+=1;
    }
    function get_v() public view returns (uint)
    {
        return v;
    }
    function get_candidate(uint candidate_num) public view returns (uint)
    {
        return candidate[candidate_num];
    }
    /*
    function lock() public {
        while(m == 1){}
        m=1;
    }
    function unlock() public {
        m=0;
    }
    */
}
