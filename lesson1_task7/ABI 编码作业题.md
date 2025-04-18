# ABI 编码作业题

可参考ABI应用二进制接口说明链接：https://learnblockchain.cn/docs/solidity/abi-spec.html



### ABI 编码中的 `uint<M>` 表示什么？

uint<M>：M位的无符号整数类型，`0 < M <= 256`，`M % 8 == 0`。例如 `uint32`，`uint8`，`uint256`



### 在 ABI 中，动态类型和静态类型有什么区别？

静态类型会被直接编码，动态类型在当前块之后的单独分配位置被编码。



### 解释函数选择器(function selector)在 ABI 中的用途

函数选择器：指某个函数签名的 Keccak256 哈希的前 4 个字节

补充一下什么是函数签名：函数名（逗号分隔的参数类型)；**注意**：在函数签名中，`uint`和`int`要写为`uint256`和`int256`。

函数选择器作用：用于指定调用的具体函数。



### 在 Solidity 中，哪些类型不被 ABI 直接支持？

Solidity 中的元组类型不被 ABI 直接支持，需要特定的处理。



### 如何通过 ABI 编码调用具有多个参数的函数？

通过将所有参数的编码合并，其中静态参数直接编码，动态参数先记录偏移量然后在数据部分单独编码。



### 什么是“严格编码模式”？

严格编码模式要求编码偏移量必须尽可能小，且数据区域不允许有重叠或间隙。



### 在 ABI 中，`fixed<M>x<N>` 和 `ufixed<M>x<N>` 有何不同？

`fixed<M>x<N>` 是有符号的固定小数点数，而 `ufixed<M>x<N>` 是无符号的固定小数点数。其中 M 是总位数，N 是小数位数。



### 事件的 ABI 编码如何处理已索引和未索引的参数？

 已索引的参数将与事件的 Keccak 哈希一起作为日志项的主题存储。未索引的参数则存储在日志的数据部分。



### 描述如何通过 ABI 对一个返回错误的函数进行编码

错误函数的编码与普通函数相似，但使用错误选择器。例如，`InsufficientBalance` 错误将编码其参数并使用特定的错误选择器。



### `abi.encodePacked()` 在什么情况下使用，它与 `abi.encode()` 有何区别？

abi.encodePacked()：`abi.encodePacked()` 用于非标准打包模式，适用于需要紧凑编码的情况。会进行压缩编码，

而abi.encode()会对不足32字节的类型进行补0操作；



### 解释 ABI 中对动态数组编码的过程

动态数组首先编码数组长度，然后编码数组中每个元素。如果元素是动态类型，则对每个元素进行独立编码，并记录其偏移



### 如何在 ABI 中处理嵌套数组或结构体？

嵌套数组或结构体按其元素顺序编码，每个元素根据其类型（静态或动态）适当处理。动态元素会记录偏移量，然后编码其内容。

