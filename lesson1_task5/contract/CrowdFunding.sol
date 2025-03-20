// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

// 众筹合约
contract CrowdFunding {
    // 状态变量
    address public beneficiary;                 // 受益人
    uint public goal;                           // 目标金额
    uint public amount;                         // 已筹金额
    mapping(address => uint256) public funders; // 贡献者及其金额
    bool public funded = false;                 // 是否达到目标金额
    bool public fundingClosed = false;          // 众筹是否结束(众筹失败或众筹超期)

    constructor(address beneficiary_, uint goal_){
        beneficiary = beneficiary_;
        goal = goal_;
    }

    // 回退函数
    // fallback() external payable {
    //     contribute();
    // }

    // 资助
    function contribute() external payable {
        require(!fundingClosed, "CrowdFunding is closed");

        // 判断已筹金额是否大于目标金额
        uint256 potentialFundingAmount = amount + msg.value;
        if (potentialFundingAmount > goal) {
            uint refundAmount = potentialFundingAmount - goal; // 退款金额
            funders[msg.sender] += msg.value - refundAmount;   // 贡献金额 — 退款金额
            amount += msg.value - refundAmount;                // 记录已筹金额
            payable(msg.sender).transfer(refundAmount);  // 退还超过的金额
        } else {
            amount += msg.value;                // 记录已筹金额
            funders[msg.sender] += msg.value;   // 记录贡献者及其金额
        }
        checkGoal(); // 检查是否达到筹款目标
    }

    // 检查是否达到筹款目标
    function checkGoal() internal {
        if (amount >= goal) {
            funded = true;
            fundingClosed = true;
            // 可以在这里添加代码来处理达到目标后的逻辑，比如通知受益者等
            // ...
        }
    }

    // 受益者在达到筹款目标后提取资金
    function claimFunds() public {
        require(funded, "Funding goal not reached");
        require(msg.sender == beneficiary, "Not the beneficiary");
        payable(beneficiary).transfer(goal);
    }

    // 投资者在众筹失败后取回资金
    function withdraw() public {
        require(!fundingClosed, "CrowdFunding is closed");              // 检查众筹是否结束(结束)
        require(!funded, "Goal has been reached");                      // 检查是否达到筹款目标
        require(funders[msg.sender] > 0, "No contribution found");      // 检查是否资助过

        uint amount_to_withdraw = funders[msg.sender];
        funders[msg.sender] = 0;
        amount -= amount_to_withdraw;
        payable(msg.sender).transfer(amount_to_withdraw);         // 取回资金
    }
}
