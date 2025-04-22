package main

import "fmt"

/*
变量num1定义在main()函数中，作用域就是main()函数的函数体，无法在其他函数中使用该变量

	报错信息如下：

		.\22.函数作用域_局部变量.go:9:6: declared and not used: num1
		.\22.函数作用域_局部变量.go:14:26: undefined: num1
*/
func main() {
	// .\22.函数作用域_局部变量.go:14:6: declared and not used: num1
	//declared and not used: num1 ==》已声明但未使用：num1
	var num1 int = 1
	fmt.Printf("num1 = %d", num1)
	a()
}

func a() {
	// .\22.函数作用域_局部变量.go:19:26: undefined: num1
	// undefined: num1 ==》 未定义的变量num1

	//fmt.Printf("num1 = %d", num1) // 变量num1定义在main()函数中，作用域就是main()函数的函数体，无法在其他函数中使用该变量
}
