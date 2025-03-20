// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// WETH合约：以太币本身并不符合ERC20标准。WETH的开发是为了提高区块链之间的互操作性
// WETH (Wrapped ETH)是ETH的带包装版本。我们常见的WETH，WBTC，WBNB，都是带包装的原生代币
contract Weth is ERC20("Wrapper eth", "WETH") {
    // 事件：存款和取款
    event Deposit(address indexed to, uint256 amount);
    event Withdraw(address indexed from, uint256 amount);

    // 2种方式给父级合约的构造函数设置值 
    // contract Weth is ERC20("Wrapper eth", "WETH") {}
    // constructor() ERC20("Wrapper eth", "WETH") {}

    receive() external payable {
        desposit();
    }

    fallback() external payable {
        desposit();
    }

    // 存款
    function desposit() public payable {
        _mint(msg.sender, msg.value);   // 铸币
        emit Deposit(msg.sender, msg.value);
    }

    // 取款
    function withDraw(uint256 _amount) public {
        require(balanceOf(msg.sender) >= _amount);
        _burn(msg.sender, _amount);    // 烧币
        payable(msg.sender).transfer(_amount);
    }
}
