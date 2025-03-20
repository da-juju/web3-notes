// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

// 多签钱包概念：特点是交易被多个私钥持有者（多签人）授权后才能执行：
// 例如钱包由3个多签人管理，每笔交易需要至少2人签名授权。
// 优点：多签钱包可以防止单点故障（私钥丢失，单人作恶），更加去中心化，更加安全，被很多DAO采用。
contract MultiSigWallet {
    // 收款事件
    event Deposit(address indexed sender, uint256 amount);
    // 提交交易单
    event Submit(uint256 indexed txId);
    // 多签人批准交易单
    event Approve(address indexed owner, uint256 indexed txId);
    // 多签人在交易单未执行的情况下，可撤销审批
    event Revoke(address indexed owner, uint256 indexed txId);
    // 执行交易，将主币发送到目标地址
    event Execute(uint256 indexed txId);

    // -----------------------------------------------------------

    // 交易单结构体：转账接受人address，交易金额，data，交易状态（是否执行）
    struct Transaction {
        address to; // 转账接受人address
        uint256 amount; // 交易金额
        bytes data; // 如果交易收款人地址是一个合约地址，可以调用合约的函数，data就是合约函数的参数值
        bool executed; // 交易状态（表示：该交易是否已执行）
    }

    // owners：多签持有人数组
    // isOwner：address => bool的映射，记录一个地址是否为多签持有人
    // threshold：多签执行门槛，交易至少有n个多签人签名才能被执行。
    address[] public owners;
    mapping(address => bool) public isOwner;
    uint256 public threshold;

    Transaction[] public transactions;
    // txId -> owner -> true：某一笔交易，是否有多签人已批准
    mapping(uint256 => mapping(address => bool)) public approved;

    // -----------------------------------------------------------
    // 是否为合约拥有者
    modifier onlyOwner() {
        require(isOwner[msg.sender], "not owner");
        _;
    }

    // -----------------------------------------------------------

    // 初始化多签人和多签最少确认数量
    // 需要添加校验
    constructor(address[] memory _owners, uint256 _threshold) {
        require(_owners.length > 0, "owners required");
        require(
            _threshold > 0 && _threshold <= _owners.length,
            "threshold should be less than or equal number of owners"
        );

        // 校验多签人地址是否有效
        // 校验多签人地址是否重复
        for (uint256 i = 0; i < _owners.length; i++) {
            address owner = _owners[i];
            require(owner != address(0), "invalid owner address");
            require(!isOwner[owner], "repeated owner address");
            owners.push(owner);
            isOwner[owner] = true;
        }
        threshold = _threshold;
    }

    // 收款函数
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    // 提交交易单
    function submit(
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external onlyOwner {
        transactions.push(
            Transaction({
                to: _to,
                amount: _amount,
                data: _data,
                executed: false
            })
        );
        emit Submit(transactions.length - 1); // txId = transactions数组长度减1（也就是数组的下标）
    }

    // 审批交易单
    // 校验：是否为拥有者、交易id是否存在、交易是否已被审批、交易是否已执行
    function approve(uint256 _txId) external onlyOwner {
        approved[_txId][msg.sender] = true;
        emit Approve(msg.sender, _txId);
    }

    // 撤销交易
    function revoke(uint256 _txId) external onlyOwner {
        approved[_txId][msg.sender] = false;
        emit Revoke(msg.sender, _txId);
    }

    // 执行交易
    function execute(uint256 _txId) external onlyOwner {
        // 如何获取交易信息。交易单存储在transactions数组时，txId对应着存储的下标index ==》可通过下标获取交易单
        Transaction storage currentTx = transactions[_txId];
        uint256 approvedCount = _getApproveCount(_txId);
        require(
            approvedCount >= threshold,
            "approved number less than threshold"
        ); // 校验是否满足多签执行门槛
        // 1、当前交易单执行状态变更
        currentTx.executed = true;
        // 2、发送主币
        // call{value: 123}("")
        /// {value: 123}：指转账金额
        /// ("")：指发送的数据
        (bool success,) = currentTx.to.call{value: currentTx.amount}(
            currentTx.data
        );
        require(success, "Transaction execution failed");
        emit Execute(_txId);
    }

    // 私有函数：获取当前交易单的审批数量
    function _getApproveCount(uint256 _txId) private view returns (uint256) {
        uint256 count;
        for (uint256 i = 0; i < owners.length; i++) {
            if (approved[_txId][owners[i]]) {
                count += 1;
            }
        }
        return count;
    }
}
