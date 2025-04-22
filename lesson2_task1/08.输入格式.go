package main

import "fmt"

/*
接受键盘输入值（类似java Scanner）,Go提供了2种接受键盘输入的方式

	1.Scan()
	2.Scanf()
*/
func main() {

	// fmt.Scanf()
	var age int
	fmt.Println("请输入用户年龄：")
	// 将键盘输入的值，保存到变量age中
	fmt.Scanf("%d", &age) //&age: "&" => 表示获取age变量的内存地址
	fmt.Println("age = ", age)
	// 获取age变量的内存地址
	fmt.Println("age的内存地址", &age) //0x140a0e8
	// %p:内存地址格式符
	fmt.Printf("age的内存变量地址：%p\n", &age)

	// fmt.Scan()
	fmt.Println()
	fmt.Println("请重新输入用户年龄：")
	fmt.Scan(&age)
	fmt.Println("age = ", age)
}
