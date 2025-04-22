package main

import "fmt"

/*
	方法值（method value）和方法表达式（method expression）是两种与接收者（receiver）相关的重要概念。
	它们允许你将方法绑定到特定的实例（或改变方法的接收者类型），然后像普通函数一样调用或传递它们。

方法值和方法表达式的方式，实现方法的调用

 1. 方法值

 2. 方法表达式
*/
type User struct {
	name string
	age  int
}

func (u *User) PrintInfo() {
	fmt.Println("userInfo=", *u)
}

func main() {
	var user = User{"张三", 18}

	// 对象调用方法
	user.PrintInfo()

	// 1.方法值
	f := user.PrintInfo
	fmt.Println("f=", f)        // 0xc78850
	fmt.Printf("f变量类型=%T\n", f) // func()
	f()

	// 2.方法表达式
	f1 := (*User).PrintInfo
	fmt.Println("f1=", f1)        // 0xec8640
	fmt.Printf("f1变量类型=%T\n", f1) // func(*main.User)
	f1(&user)
}
