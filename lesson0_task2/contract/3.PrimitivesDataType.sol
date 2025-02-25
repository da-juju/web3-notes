// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Primitives {
    /*
        基本数据类型：bool、uint、int、address
    */
    
    /* 
        uint：无符号整数，非负数
        uint8    0 ~ 2**8-1
        uint16   0 ~ 2 ** 16-1
        ...
        uint256  0 ~ 2 ** 256-1
    */
    uint num0 = 1;   // uint是uint256的别名(uint表示uint256类型)

    /*
        int：允许负数
        int8：-2**7 ~ 2**7-1
        int256：-2**255 ~ 2**255-1
    */
    int num2 = -1;

    // int的最大值和最小值
    int minnum = type(int).min;
    int maxnum = type(int).max;

    // adress:地址类型
    address public addr = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;

    // 默认值：通常用于存储以太坊账户的地址
    bool public defBoo = false;
    uint public  defUint = 0;
    int public defInt = 0;
    address public defAddr =  0x0000000000000000000000000000000000000000;   // 40个十六进制数字0组成

}