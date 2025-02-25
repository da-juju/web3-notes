// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/*
    solidity中，变量有3种
    局部变量：在函数内部声明，不存储在区块链上
    状态变量：在函数外部声明，存储在区块链上
    全局变量：获取区块链的相关的信息

*/
contract Variables {

    uint public i;    // 状态变量

 function doSomething() public {
        // Local variables are not saved to the blockchain.
        uint256 i = 456;

        // Here are some global variables
        uint256 timestamp = block.timestamp; // Current block timestamp
        address sender = msg.sender; // address of the caller
    }
    
}