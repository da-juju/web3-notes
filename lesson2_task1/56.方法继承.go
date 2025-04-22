package main

import "fmt"

/*
《方法继承》
	子类继承父类，同时可以继承父级的成员以及父类的方法
*/

type Son struct {
	Parent // 匿名字段继承
	score  float64
}

type Parent struct {
	id   int
	name string
	age  int
}

// PrintParentInfo 为父类创建打印的方法
func (p *Parent) PrintParentInfo() {
	fmt.Println("parentInfo=", *p)
}

func main() {
	son1 := Son{Parent{101, "张三", 18}, 88.8}

	// 子类继承了父类的方法，可以打印父类的信息
	son1.PrintParentInfo()
}
