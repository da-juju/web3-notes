// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract MyMapping {
    // 定义一个简单映射
    mapping(string => bool) map;

    function exapple1(string memory _name) external returns (bool) {
        // 赋值
        map["zhangsan"] = true;
        map["lisi"] = true;
        map["wangwu"] = false;
        bool b = map[_name];
        return b;
    }
}
