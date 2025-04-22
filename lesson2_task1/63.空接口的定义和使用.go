package main

import "fmt"

/*
《空接口》
	空接口（interface{}）是一种特殊的接口，它不包含任何方法。
	由于其不包含任何方法，TODO 因此任何类型都隐式地实现了空接口（反之：空接口类型可以存储任意类型数据）。
		空接口的这一特性使其变得非常灵活和强大，经常用于以下场景：

一、空接口的定义
	空接口的定义非常简单，直接使用 interface{} 即可：
		var i interface{}
		这里的 i 就是一个空接口类型的变量

二、空接口的使用
	1.作为任意类型的容器：
		由于任何类型都实现了空接口，因此空接口可以存储任何类型的值。
	2.类型断言和类型选择：
		当我们从空接口中取出值时，通常需要使用类型断言或类型选择来确定其实际类型。
*/

func printValue(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type is int:", v)
	case string:
		fmt.Println("Type is string:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {

	// 一、空接口的定义

	// 二、空接口的使用
	// 1.作为任意类型的容器：
	//var box interface{}
	//box = 42
	//box = "hello"
	//box = true

	//2.类型断言和类型选择：
	printValue(1)
}
