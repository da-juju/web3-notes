package main

import "fmt"

type Animal struct {
	name string
}

type Dog struct {
	*Animal // 指针类型匿名字段继承父类
	breed   string
}

func main() {

	dog1 := Dog{&Animal{"大黄"}, "狗"}
	fmt.Println("dog1=", dog1) // dog1= {0x1508078 狗}

	// 指针类型匿名字段继承类，不能直接打印。只能通过成员变量的方式访问
	fmt.Println("name=", dog1.name)   // 大黄
	fmt.Println("breed=", dog1.breed) // 狗
}
