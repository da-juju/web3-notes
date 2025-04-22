package main

import "fmt"

func main() {
	var name string = "itcast"
	fmt.Printf("%s\n", name) // itcast

	str := "a"
	// %T：打印变量类型
	fmt.Printf("%T\n", str) // string
	fmt.Printf("%T\n", 123) // int

	// 字符串类型
	// 注意：虽然看到的str中只是包含了一个字符串，但是隐藏者一个字符串的结束标记 '\0'

	// 字符串中的字符个数
	fmt.Println(len(name)) // "itcast"字符串中的字符个数为：6

	var name2 string = "中国人"
	fmt.Println(len(name2)) // 9 ==> 一个中文汉族占用3个字符

}
