// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract IfElse {
    // 判断是否成年
    function foo(uint256 _age) external pure returns (bool) {
        if (_age >= 18) {
            return true;
        } else {
            return false;
        }
    }

    function foo2(uint256 _age) external pure returns (bool) {
        //   return   _age > 18 ? true : false;
        return _age >= 18;
    }
}
