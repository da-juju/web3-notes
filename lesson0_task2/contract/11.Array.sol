// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


/*
    数组：引用类型，可存储一组数据（整数，字节，地址等等）
    格式：定长数组：T[V] T表示类型，V表示数组长度
          动态数组：T[] T表示类型
    初始化和赋值:
        定长数组：
            uint[5] fixedArray = [1, 2, 3, 4, 5];
        动态数组：
            uint[] dynamicArray = [1, 2, 3]，或者dynamicArray.push(4)来添加元素
*/      
contract Array {

        uint[5] public fixedArray = [1, 2, 3, 4, 5];
        uint[] public dynamicArray;



        // 删除数组元素
        // _index = 1
        // 1, 2, 3, 4, 5 ==> 1,0,3,4,5
        function  deleteFixedArray(uint _index) external  {
            delete  fixedArray[_index];
        }

        // 查看定长数组元素 
        function getFixedArray() external view returns (uint[5] memory){
            return fixedArray;
        }

        // 动态数组添加元素
         function addDynamicArray(uint _item) external {
            dynamicArray.push(_item);
        }

        // 移除元素（尾部移除）
        function removeDynamicArray() external {
            dynamicArray.pop();
        }

        // 查看定长数组元素 
        function getDynamicArray() external view returns (uint[] memory){
            return dynamicArray;
        }
    
}