// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract WhileDemo {

    function demo1() external pure returns (uint) {
        uint i = 0;
        while (i < 10) 
        {
            i ++;
        }
        return i;
    }
}