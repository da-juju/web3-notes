package main

import "fmt"

// 申明结构体
type Student struct {
	id   int
	name string
	age  int
	addr string
}

/*
1.结构体切片定义

	var sliceVar = []切片类型{...}

2.修改结构体成员变量

	sliceVar[下标].成员 = 值

3.循环遍历
*/
func main() {

	fmt.Println("//  1.结构体切片定义  //")
	// 1.结构体切片定义
	var studentSlice = []Student{
		{id: 1, name: "张三", age: 18, addr: "滨湖区"},
		{id: 2, name: "李四", age: 30, addr: "新吴区"},
		{id: 3, name: "老王", age: 44, addr: "锡山区"},
	}
	fmt.Println(studentSlice)

	fmt.Println("//  2.修改结构体成员变量  //")
	// 2.修改结构体成员变量
	// sliceVar[下标].成员 = 值
	studentSlice[0].age = 55
	fmt.Println(studentSlice)

	fmt.Println("//  3.循环遍历  //")
	// 3.循环遍历
	for _, student := range studentSlice {
		fmt.Println("-------------------")
		fmt.Println(student)
		fmt.Println(student.name)
	}
}
