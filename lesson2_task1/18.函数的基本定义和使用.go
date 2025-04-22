package main

import "fmt"

/*
一、函数定义

	在 Go 语言中，函数是一种基本的代码构造块，【todo 用于封装可重复使用的代码段】。
	Go 函数具有明确的类型签名，可以接收参数并返回结果。以下是关于 Go 函数的一些关键点：

二、基本申明语法

		函数通过 func 关键字定义，其基本语法如下：
		func functionName(parameters) (results) {
	    	// 函数体
		}

		# functionName 是函数的名称。
		# parameters 是函数的输入参数，用括号括起来，参数之间用逗号分隔。如果函数没有参数，括号内可以为空。
		# results 是函数的返回值，也用括号括起来。如果函数没有返回值，可以省略返回值部分或显式地声明为 ()。
		# 函数体由大括号 {} 包围，包含执行的具体代码。

三、函数调用

	functionName(parameters)
*/
func main() {
	// 函数调用
	playGame()
	play()
	playGame()
}

// 函数申明
func playGame() {
	fmt.Println("超级玛丽，走啊走，跳啊跳")
	fmt.Println("超级玛丽，走啊走，跳啊跳")
}

func play() {
	fmt.Println("屏幕闪烁")
	fmt.Println("播放音乐")
}
