package main

import "fmt"

/*
《recover错误拦截》

	recover 是一个内置函数，它用于从 panic 引发的异常中恢复程序的正常运行。
	TODO 注意：recover函数只有的在defer调用的函数中有效


	在上面的例子中，尽管 causePanic 函数引发了 panic，
	但 safeCall 函数中的 defer 注册的匿名函数捕获了这个恐慌，并通过 recover 恢复了程序的执行。
	因此，fmt.Println("已恢复：", r) 会被执行，而 fmt.Println("这行代码不会被执行") 则不会。
*/
func main() {
	safeCall()
}

func causePanic() {
	panic("发生了恐慌！")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("已恢复：", r)
		}
	}()
	causePanic()
	fmt.Println("这行代码不会被执行")
}
