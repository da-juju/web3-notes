package main

import "fmt"

/*
一、概念
for range 循环

	for range 循环是Go语言中一种更简洁的迭代方式，它会自动处理索引和值。

二、基本语法

			// index：集合的索引位置
			// value：集合元素
			for index, value := range collection {
		    	// 使用 index 和 value 进行操作
			}

			// 如果不需要索引，可以使用下划线 _ 忽略它：
			for _, value := range collection {
	    		// 使用 value 进行操作
			}
*/
func main() {
	// 使用短变量声明和初始化
	arr := [5]int{1, 2, 3, 4, 5}
	for index, el := range arr {
		fmt.Printf("索引：%d ==> 元素：%d\n", index, el)
	}

	fmt.Println()

	// 如果不需要索引，可以使用下划线 _ 忽略它：
	for _, value := range arr {
		// 使用 value 进行操作
		fmt.Printf("==> 元素：%d\n", value)
	}
}
