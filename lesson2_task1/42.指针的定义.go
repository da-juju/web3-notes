package main

import "fmt"

/*
一、概念

	指针也是一种特殊的变量，用于存储变量的内存地址。通过使用指针，你可以直接操作内存中的数据

二、声明和初始化指针

	指针通过 *Type 形式声明，其中 Type 是指针指向变量的类型。
		示例：
			var p *int       // 指针变量申明
			var a int = 42
			p = &a           // 获取变量 a 的地址并赋值给指针 p

三、获取指针对应的变量值

	1.获取指针的变量值：*指针变量
	2.修改指针的变量值：*指针变量 = 新值

四、常用方法

 1. &arg：获取arg变量的内存地址
 2. *p：获取p指针对应的变量值(todo *p：指针解引用 ==> 是指通过指针访问它所指向的变量的值)
*/
func main() {

	var a int = 10
	fmt.Println("a=", a)

	// a变量对应的指针
	var p *int
	p = &a
	fmt.Println("&a=", p)

	// 获取p指针对应的变量值
	var b = *p
	fmt.Println("b=", b)

	// 通过指针操作内存中的数据
	*p = 222
	fmt.Println("*p=222：", *p)
	fmt.Println("a=222：", a)

}
