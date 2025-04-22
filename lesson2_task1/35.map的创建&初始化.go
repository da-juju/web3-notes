package main

import "fmt"

/*
一、map概念

	map 是一种内置的数据结构，用于存储键值对（key-value pairs）。
	map 提供了非常高效的数据查找、插入和删除操作，其内部实现通常基于哈希表。

二、map创建

		1.var
			var myMap map[key类型]value类型
		2.字面量方式
			myMap :=map[key类型]value类型{}
	    3.make函数
			var myMap = make(map[key类型]value类型)
*/
func main() {

	fmt.Println("//   2.字面量方式  //")
	m1 := map[string]int{"zs": 18, "lisi": 28}
	fmt.Println("2.字面量方式申明并初始化map=", m1)

	fmt.Println("//   3.make函数  //")
	m2 := make(map[string]int)
	m2["Apple"] = 8888
	m2["Android"] = 9999
	fmt.Println("3.make函数申明并初始化map=", m2)
}
