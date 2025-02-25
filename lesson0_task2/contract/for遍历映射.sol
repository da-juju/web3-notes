// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract MappingTraversal {
    mapping(uint => string) public nameToAge;
    uint[] public keys;

    // 添加一个键值对，并记录键
    function addPerson(uint _id, string memory _name, uint _age) public {
        nameToAge[_id] = string(abi.encodePacked(_name, " ", _age));
        keys.push(_id);
    }

    // 遍历mapping并返回所有记录
    function getAllPersons() public view returns (string[] memory) {
        string[] memory result = new string[](keys.length);
        for (uint i = 0; i < keys.length; i++) {
            result[i] = nameToAge[keys[i]];
        }
        return result;
    }
}