// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


interface IERC20 {
    // 转账事件
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    // 授权事件
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    // 代币总供给量
    function totalSupply() external  view returns (uint256);
    // 查询账户余额
    function balanceOf(address _owner) external  view returns (uint256);
    // 转账
    function transfer(address _to, uint256 _value) external  returns (bool);
    // 授权
    function approve(address _spender, uint256 _value) external  returns (bool);
    // 授权转账额度
    function allowance(address _owner, address _spender) external  view returns (uint256);
    // 授权转账
    function transferFrom(address _from, address _to, uint256 _value) external  returns (bool);

}


contract MyERC20 is IERC20 {

    string public name;   // 名称
    string public symbol;  // 代号
    uint8 public decimals = 18; // 小数位数(精度)

    mapping(address => uint256) public balanceOf;   // 账户余额
    mapping(address => mapping(address => uint256)) public allowance; // 授权额度
    uint256 public totalSupply;   // 代币总供给

    constructor(string memory _name,string memory _symbol) {
        name = _name;
        symbol = _symbol;
        // 铸币10000
        _mint(10000);
    }

    // 代币总供给量（状态变量提供了同名方法）
    // function totalSupply() external  view returns (uint256){
    //     return totalSupply;
    // }

    // 查询账户余额（状态变量提供了同名方法）
    // function balanceOf(address _owner) external  view returns (uint256){
    //     return balanceOf[_owner];
    // }

    // 转账
    function transfer(address _to, uint256 _value) external  returns (bool){
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }
    // 授权
    function approve(address _spender, uint256 _value) external  returns (bool){
        allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }
    // 查询授权转账额度（状态变量提供了同名方法）
    // function allowance(address _owner, address _spender) external  view returns (uint256){
    //     return allowance[_owner][_spender];
    // }

    // 授权转账（被授权的人调用该函数）
    function transferFrom(address _from, address _to, uint256 _value) external  returns (bool){
        allowance[_from][msg.sender] -= _value;
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(_from, _to, _value);
        return true;
    }

    // 铸币函数
    function _mint(uint _total) public  {
        totalSupply += _total;
        balanceOf[msg.sender] += _total;
        emit Transfer(address(0), msg.sender, _total);
    }

    // 销毁代币
    function _burn(uint _amount) external {
        totalSupply -= _amount;
        balanceOf[msg.sender] -= _amount;
        emit Transfer( msg.sender, address(0),_amount);
    }

}
