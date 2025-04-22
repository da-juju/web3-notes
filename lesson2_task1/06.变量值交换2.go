package main

import "fmt"

func main() {
	num1 := 1
	num2 := 10

	// 使用多重赋值的方式交换num1和num2的值
	num1, num2 = num2, num1
	fmt.Println(num1, num2)
}
