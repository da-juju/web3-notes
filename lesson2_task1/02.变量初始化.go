package main

import "fmt"

/*
变量初始化
*/
func main() {

	// 方式一：申明并初始化
	var age int = 123
	fmt.Println(age)

	// 方式二，申明后初始化
	var age2 int
	age2 = 11
	fmt.Println(age2)
}
