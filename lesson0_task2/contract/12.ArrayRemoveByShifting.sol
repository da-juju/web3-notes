// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/*
    删除数组中指定下标元素
    方式一：通过从右向左移动元素来删除数组元素，然后使用pop()的方式移除尾部元素(有序，消耗Gas)
    方式二：通过替换的方式，替换尾部的元素和指定下标的元素，然后使用pop()的方式移除尾部元素（无序）
*/

contract ArrayRemove {
    uint256[] public darr = [1, 2, 3, 4, 5];
    uint256[] public darr2 = [1, 2, 3, 4, 5, 6, 7];

    // 方式一：通过从右向左移动元素来删除数组元素，然后使用pop()的方式移除尾部元素(有序，消耗Gas)
    function removeItem1(uint256 _index) external {
        uint256 len = darr.length;

        // 条件为false，提示：下标越界
        require(_index < len, "index out of bound");

        for (uint256 i = _index; i < len - 1; i++) {
            darr[i] = darr[i + 1];
        }
        darr.pop();
    }

    function get1() external view returns (uint[] memory) {
        return darr;
    }

    // 方式二：通过替换的方式，替换尾部的元素和指定下标的元素，然后使用pop()的方式移除尾部元素（无序）
    function removeItem2( uint _index) external {
         // 条件为false，提示：下标越界
        require(_index < darr2.length, "index out of bound");
        darr2[_index] = darr2[darr2.length-1];
        darr2.pop();
    }

        function get2() external view returns (uint[] memory) {
        return darr2;
    }
}
