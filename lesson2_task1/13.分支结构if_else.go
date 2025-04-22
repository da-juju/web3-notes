package main

import "fmt"

func main() {

	// 年龄
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你已经成年了！")
	} else {
		fmt.Println("你未成年")
	}
}
