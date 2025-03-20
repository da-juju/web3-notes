// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

contract MyToken is ERC20 {


    using Math for uint256;

    function max(uint256 a,uint256 b) external pure returns (uint256) {
        return a.max(b);
    }

    // 构造函数，在合约创建时自动调用
    constructor() ERC20("MyToken", "MTK") {  // MyToken：代币名称， MTK：代币简称
        // 使用 _mint 函数向合约创建者（msg.sender）发行 1000000 个代币
        // 1000000 ：代币发行量
        // decimals()：代币精度
        // 10 ** decimals() 是为了确保代币数量正确，因为 ERC20 的 decimals() 函数返回代币的小数位数
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
}