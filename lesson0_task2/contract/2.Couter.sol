// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Couter {
    uint public count;

    function increment() external  {
        count += 1;
    }

    function sub() external {
        count -= 1;
    }
}