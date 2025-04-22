package main

import "fmt"

/*
Printf：输出格式化
*/
func main() {

	num1, num2, num3 := 10, 20, 30

	// %d:表示整型变量值的格式符
	fmt.Printf("num1 = %d, num2 = %d, num3 = %d", num1, num2, num3)
	fmt.Println()
	// \n：表示换行
	fmt.Printf("num1 = %d\nnum2 = %d\nnum3 = %d\n", num1, num2, num3)
}
