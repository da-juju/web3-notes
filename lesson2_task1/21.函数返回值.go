package main

import "fmt"

func main() {
	/*
		var a int
		var b int
		fmt.Println("请输一个整型数值：")
		fmt.Scan(&a)
		fmt.Println("请再次输一个整型数值：")
		fmt.Scan(&b)

		var sum int
		// 两个整型数值相加
		sum = addResult(a, b)
		// 判断数值是否为偶数
		flag := isEvenNumber(sum)
		if flag {
			fmt.Printf("数值求和=%d,为偶数\n", sum)
		} else {
			fmt.Printf("数值求和=%d,为奇数\n", sum)
		}
	*/

	//    函数返回多个值    //
	fmt.Println("//    函数返回多个值    //")

	var x int
	var y int
	x, y = getResult()
	fmt.Printf("getResult()函数返回值：%d ==> %d", x, y)
}

// 函数返回值方式一：两个整型数值相加
func addResult(num1 int, num2 int) int {
	return num1 + num2
}

// 函数返回值方式二：两个整型数值相加
func addResult2(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

// 函数返回值方式三：两个整型数值相加
func addResult3(num1 int, num2 int) (result int) {
	result = num1 + num2
	return // 如果已经指定了返回值的变量名称，return关键字后面可以不需要添加变量名称
}

// 判断数值是否为偶数
func isEvenNumber(arg int) bool {
	if arg%2 == 0 {
		return true
	} else {
		return false
	}
}

//    函数返回多个值    //

func getResult() (num1 int, num2 int) {
	num1 = 1
	num2 = 2
	return num1, num2
}
