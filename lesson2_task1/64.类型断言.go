package main

import "fmt"

/*
《空接口类型断言》
一、作用
	通过类型断言，可以判断空接口中存储数据的类型

二、语法：
	value,ok := m.(T)
		m：空接口类型的变量
		T：断言的类型
		value：变量m的值（ok=true）
		ok：是否断言成功（true：成功，false：失败）
*/

func main() {
	// 空接口类型变量i
	var i interface{}
	i = 123

	value, ok := i.(string)
	fmt.Println("ok=", ok)
	fmt.Println("value=", value)
}
