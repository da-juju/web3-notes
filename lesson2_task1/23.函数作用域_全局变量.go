package main

import "fmt"

/*
	1、全局变量：指在任何函数中都可以访问得到得变量，并可对变量值进行修改
	2、全局变量和局部变量的优先级：局部 > 全局
*/

var count int

func main() {
	/*
		count = 10
		aa()
		fmt.Printf("count变量的值：%d", count)
	*/

	//  	2、全局变量和局部变量的优先级：局部 > 全局   //
	fmt.Println("测试 全局变量和局部变量的优先级")

	count = 10
	aa()
	fmt.Printf("全局变量{count}：%d", count)

	// 输出结果，调用aa()函数后，全局变量值没有被修改为：-1
	// todo 证明 优先级：局部变量 > 全局变量
	// todo 如果全量变量和局部变量的名字相同，优先使用的是局部变量
	//局部变量{count}： -1
	//全局变量{count}：10
}

// a
func aa() {
	/*
		count = -1
	*/
	var count int
	count = -1
	fmt.Println("局部变量{count}：", count)
}
