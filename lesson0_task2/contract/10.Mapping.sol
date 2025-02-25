// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/*
    映射：引用类型，存储key-value键值对
    定义格式：mapping(k=>v) mappingName
*/
contract Mapping {
    // 定义一个简单映射
    mapping(string => bool) public map;
    // 复杂的映射：维护朋友关系
    mapping(string => mapping(string => bool)) public isFriend;

    function isExists(string memory _name) external returns (bool) {
        // 赋值
        map["zhangsan"] = true;
        map["lisi"] = true;
        map["wangwu"] = false;
        bool b = map[_name];
        return b;
    }

    function isFriends(string memory _name1, string memory _name2)
        external
        returns (bool)
    {
        isFriend["zs"]["ls"] = true;
        isFriend["zs"]["ww"] = true;
        isFriend["ls"]["ww"] = true;
        return isFriend[_name1][_name2];
    }

    // 删除map对应的key
    function deleteMap(string memory _name) external {
        delete map[_name];
    }
}
