package main

import "fmt"

/*
一、概念

	结构体类似于java中的对象
	结构体是由一系列相同或者不同类型的数据，构成的数据集合。可以很好的管理一批有关联的数据，可以提高程序的易读性

二、申明结构体

	基本语法格式：
		type MyStruct struct {
			id int
			name string
		}
*/
func main() {

	// 申明结构体
	type Student struct {
		id   int
		name string
		age  int
		addr string
	}

	// 初始化
	// 1.顺序初始化
	var stu1 = Student{1, "张三", 18, "无锡滨湖区"}
	fmt.Println("1.顺序初始化:", stu1)

	// 2.选择性初始化
	var stu2 = Student{id: 1, age: 22, name: "李四"}
	fmt.Println("2.选择性初始化:", stu2)

	// 3."结构体.成员"的方式初始化
	var stu3 Student
	stu3.id = 3
	stu3.name = "王五"
	fmt.Println("3.\"结构体.成员\"的方式初始化:", stu3)

}
