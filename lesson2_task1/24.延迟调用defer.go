package main

import "fmt"

func main() {

	// defer关键字，可以实现延迟调用的功能
	defer fmt.Println("后执行...")
	fmt.Println("先执行..")

	//  fmt.Println("后执行...")
	//	fmt.Println("先执行..")
	// 正常执行结果：后执行... ——> 先执行...

	//defer fmt.Println("后执行...")
	//fmt.Println("先执行..")
	// 加了defer关键字
	// 执行结果：先执行... ——>  后执行...

	//  多个defer执行顺序   //
	fmt.Println()

	// 多个defer执行顺序
	// todo 先进后出 （有个压栈的过程，类似子弹夹：先放入的子弹，会被压入到弹夹的底部）
	defer fmt.Println("1...")
	defer fmt.Println("2...")
	defer fmt.Println("3...")

}
