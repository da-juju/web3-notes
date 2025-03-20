// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

// 小猪存钱罐实战
// 需求1：可存钱
// 需求2：可取钱（当合约自毁的情况下，可取钱）
// 需求3：查询余额
contract Bank {
    // 合约拥有者
    address public owner;   //0x5B38Da6a701c568545dCfcB03FcB875f56beddC4

    // 存款事件
    event Deposit(uint256 amount);
    // 取款事件
    event Withdraw(uint256 amount);

    constructor() {
        owner = msg.sender;
    }

    // 存款:通过回退函数接受主币
    receive() external payable {
        emit Deposit(msg.value);
    }

    // 取款:合约拥有者可取
    function withdraw() external {
        require(owner == msg.sender, "not owner");
        emit Deposit(address(this).balance);
        selfdestruct(payable(msg.sender)); // 销毁合约，发送主币给合约拥有者
    }
}
