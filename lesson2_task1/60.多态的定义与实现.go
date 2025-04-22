package main

import "fmt"

/*
《多态》
一、多态定义
	多态：指一个接口的多种表现形式。（同一个接口，使用不同的实例执行不同的操作）

二、多态的实现方式
	// 实现多态特性的函数：根据 参数传入不同类型的接口的实现类，可以完成不同的操作
	func 函数名 (参数 接口类型){
		......
	}

三、代码案例
		如下Speaker接口定义了一个Speak()的模板方法。然后Person和Dog2实现了接口方法
		// 定义一个公共的Speak方法
		func commonSpeak(s Speaker) string {
			return s.Speak()
		}

		// commonSpeak()方法传入不同的实例，输出不同的结果。这种现象就是多态的体现
		fmt.Println(commonSpeak(p)) // Hello, my name is Alice
		fmt.Println(commonSpeak(d)) // Woof!

*/

// Speaker 定义一个接口
type Speaker interface {
	Speak() string
}

// Person 定义一个类型，并实现 Speaker 接口
type Person struct {
	name string
}

// Speak 为 Person 类型实现 Speak 方法
func (p Person) Speak() string {
	return "Hello, my name is " + p.name
}

// Dog2 另一个类型，也实现 Speaker 接口
type Dog2 struct{}

// Speak 为 Dog 类型实现 Speak 方法
func (d Dog2) Speak() string {
	return "Woof!"
}

// 定义一个公共的Speak方法
func commonSpeak(s Speaker) string {
	return s.Speak()
}

func main() {
	var p = Person{name: "Alice"}
	// 使用 Person 类型实例
	fmt.Println(commonSpeak(p)) // Hello, my name is Alice
	// 使用 Dog 类型实例
	d := Dog2{}
	fmt.Println(commonSpeak(d)) // Woof!

}
