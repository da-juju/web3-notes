package main

import "fmt"

/*
指针数组
一、概念

	数组的元素是指针类型，也就是一个存储了地址的数组

	var 数组名 [下标]*类型
*/
func main() {

	// 定义指针数组
	var arr [2]*int
	var a = 1
	var b = 2

	// 指针数组初始化
	arr[0] = &a
	arr[1] = &b
	fmt.Println(arr) // [0x180a0e8 0x180a0ec]

	// 根据指针数组解引用数组元素
	fmt.Println(*arr[0]) // 1

}
