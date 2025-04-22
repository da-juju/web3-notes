package main

import "fmt"

/*
在Go语言中，异常处理主要通过内置的panic和recover机制来实现。
		1. panic
			# panic是一个内置函数，当程序遇到无法恢复的错误时，可以调用它使程序进入恐慌状态。
			# 当panic被调用时，程序的正常执行流程会立即中断，并开始逐层向上执行已注册的延迟（deferred）函数。
			# 如果在panic后没有通过recover恢复，程序会终止并打印出传递给panic的值。
		2. recover
			# recover是一个内置函数，用于从panic引发的恐慌状态中恢复。
			# 它只能在延迟函数中调用，并且只有在panic引发的恐慌状态正在传播时，recover才能捕获到恐慌值。
			# 如果recover成功捕获到恐慌值，那么程序会恢复正常的执行流程，并从调用panic的函数返回。
			# recover的返回值是传递给panic的值，如果没有恐慌状态正在传播，recover会返回nil。
*/

func main() {
	testPanic()
}

func testPanic() {
	fmt.Println("aaa")
	// 2.程序中出现严重异常，系统内部会调用panic()函数，终止程序的执行
	panic("进入恐慌状态") // 1.引发异常，强制终止整个程序执行
	fmt.Println("hello")
}
