// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/*
    常量：在Solidity中，constant和immutable都是用来定义不可更改的值的
*/

contract Constants {
    // constant:声明时必须初始化
    string public constant s1 = "constant";

    // immutable：可以先声明后初始化，可以在构造函数种初始化
    uint256 public immutable num;

    constructor(uint256 _num) {
        num = _num;
    }
}
