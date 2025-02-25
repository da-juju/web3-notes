// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


// 构造函数是在合约创建时执行的可选函数
contract  MyContract {

    string public text;
    constructor(string memory _text){
        text = _text;
    }
}

// Base contract X
contract X {
    string public textX;

    constructor(string memory _name) {
        textX = _name;
    }
}


// Base contract Y
contract Y {
    string public textY;

    constructor(string memory _text) {
        textY = _text;
    }
}

// 有两种方法可以使用参数初始化父合约

// 方式一：在继承列表中传递参数.
contract B is X("Input to X"), Y("Input to Y") {}
// 方式二：在构造函数中传递参数
contract C is X, Y {
    constructor(string memory _name, string memory _text) X(_name) Y(_text) {}
}
