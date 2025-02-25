// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Loop {

    uint[] arr1;

    function foo(uint _num) external  pure returns (uint,uint)  {
        // for循环
     // 累加器
        uint256 s;
        for (uint256 i = 1; i <= _num; i++) {
            s += i;
        }
        

        // while循环
        uint256 j;
        while (j < 10) {
            j++;
        }
        
        return (s,j);
    }
}
