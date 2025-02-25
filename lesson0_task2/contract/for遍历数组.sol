// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract ForExample {
    // for循环提供了两种方式：普通for和foreach
    uint256[] unms = [1, 2, 3, 4];

    function forDemo() external view returns (uint256[] memory) {
        uint256[] memory result = new uint256[](unms.length);
        for (uint256 i = 0; i < unms.length; i++) {
            result[i] = unms[i] * 2;
        }
        return result;
    }

    // function foreachDemo() external view returns (uint256[] memory) {
    //     uint256[] memory result = new uint256[](nums.length);
    //     uint256 index = 0;
    //     for (uint256 num : nums) {
    //         result[index++] = num * 2;
    //     }
    //     return result;
    // }
}
