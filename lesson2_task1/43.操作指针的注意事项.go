package main

import "fmt"

/*
操作指针的注意事项

	1.空指针 （nil）
	2.不能操作没有合法指向的指针（比如：不能操作空指针）
	3.new()函数开辟内存空间
*/
func main() {

	// 1.申明一个指针变量
	var p *int
	fmt.Println(p) // <nil>

	// 2.不能操作没有合法指向的指针
	// *p = 66
	// fmt.Println(*p) // invalid memory address or nil pointer dereference(无效内存地址或nil指针解引用)

	// 3.new()函数开辟内存空间
	p1 := new(int)             // 实例化，未初始化
	fmt.Println("指针变量p1=", p1) // 0x188c0bc
	fmt.Println("指针解引用 *p1=", *p1)
	*p1 = 88
	fmt.Println("指针解引用赋值 *p1=", *p1)
}
