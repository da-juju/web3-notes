package main

import "fmt"

/*
结构体指针

	结构体指针则是指向结构体的指针，它存储了结构体的内存地址。
	使用结构体指针可以带来一些好处，比如能够直接修改原始结构体的内容，以及在函数间传递大型结构体时避免复制开销。


	1.声明与初始化
		使用 * 操作符来声明一个指向该结构体的指针。
			①申明：
				var personPtr *Person
			②初始化：
				person := Person{Name: "Alice", Age: 30}
				personPtr = &person
	2.值的修改
	3.循环遍历
*/
func main() {

	// 定义结构体
	type Person struct {
		name string
		age  int
	}

	// 1.声明与初始化
	fmt.Println(" // 1.声明与初始化")
	// 申明结构体指针
	var personPtr *Person
	// 初始化
	person := Person{name: "Alice", age: 30}
	personPtr = &person

	// 2.值的修改
	fmt.Println(" // 2.值的修改")

	// 方式一：通过解引用的方式修改值
	(*personPtr).name = "Bob" // 修改结构体的字段
	fmt.Println(*personPtr)   // 输出: {Bob 30}

	// 方式二：通过指针的方式修改值
	personPtr.name = "Rose"
	personPtr.age = 22
	fmt.Println(*personPtr) // 输出: {Rose 22}

}
