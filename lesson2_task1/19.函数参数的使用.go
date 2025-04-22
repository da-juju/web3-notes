package main

import "fmt"

/*
函数的参数

	1.一个参数

		func sum(value int){...}

	2.多个参数

		func add(a int, b int){...}

	3.不定参数列表（不确定函数参数个数）

		func
*/
func main() {
	// 一个实参
	sum(100)
	sum(200)

	fmt.Println()

	// 多个实参
	add(1, 2)

	fmt.Println()

	// 不定参数列表
	multipleNumAdd(1, 1, 2)

}

// 函数定义一个形参
// 1~n之间数字求和
func sum(value int) {
	var sum = 0
	for i := 1; i <= value; i++ {
		sum += i
	}
	fmt.Printf("1~%d之间数字求和：%d\n", value, sum)
}

// 函数定义多个形参
func add(a int, b int) {
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

// 多个数字累加
func multipleNumAdd(args ...int) {
	var sum = 0
	var sum2 = 0

	// for i：通用的循环结构
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Printf("多个数字累加：%d", sum)

	fmt.Println()

	// for range：更简洁的迭代方式
	for _, val := range args {
		sum2 += val
	}
	fmt.Printf("多个数字累加：%d", sum2)
}
