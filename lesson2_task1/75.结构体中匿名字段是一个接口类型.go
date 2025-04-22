package main

import "fmt"

/*
《匿名字段是一个接口类型》

	如果匿名字段是一个接口类型，那么包含该字段的结构体将能够调用接口中定义的所有方法，
	TODO 但前提是必须为该匿名字段分配一个实现了该接口的具体类型的值。

*/

// Describer 定义一个接口
type Describer interface {
	Describe() string
}

// A 定义一个结构体 A，它实现了 Describer 接口
type A struct{}

// Describe 结构体 A，它实现了 Describer 接口
func (a A) Describe() string {
	return "I am A"
}

// B 定义一个结构体 B，它有一个匿名字段，该字段是 Describer 接口类型
type B struct {
	Describer // 匿名字段，类型为 Describer 接口
}

func main() {
	// 创建一个 A 实例，并将其赋值给 B 结构体的匿名字段（通过 B 的构造函数或直接赋值）
	var b B
	b.Describer = A{} // 为 B 的匿名字段分配一个 A 实例

	// 调用接口方法，实际上是调用了 A 的 Describe 方法
	fmt.Println(b.Describe()) // 输出: I am A

	// 另一种方式：直接创建 B 的实例并初始化匿名字段
	bb := B{Describer: A{}}
	fmt.Println(bb.Describe()) // 输出: I am A
}
