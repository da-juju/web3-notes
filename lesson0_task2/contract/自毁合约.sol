// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract SelfDestructExample {
    address payable public owner;

    constructor() {
        // address 转 payable address 类型
        owner = payable(msg.sender);
    }

    // 函数用于销毁合约并发送剩余的以太币到owner地址
    function destroy() external {
        require(msg.sender == owner, "Only the owner can destroy the contract");
        selfdestruct(owner);
    }

}
