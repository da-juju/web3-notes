package main

import "fmt"

/*
《方法的重写》
	通常指的是在子类中提供一个与父类中方法同名且签名相同的方法。
	这样，当通过子类的实例调用该方法时，将执行子类中提供的方法实现，而不是父类中的实现。

	方法重写的特点：
     	1. 子类实例调用重写方法，默认调用的是子类中的重写方法
     	2. 如果想调用父类中的方法，可以通过：子类实例.父类.方法()

*/

// 父类
type Base struct{}

func (b Base) Greet() {
	fmt.Println("Hello from Base")
}

// 子类
type Derived struct {
	Base // 匿名字段继承
}

// 在Derived中提供与Base中同名的方法，实现“重写”效果
//func (d Derived) Greet() {
//	fmt.Println("Hello from Derived")
//}

func main() {
	// 子类实例对象
	derived := Derived{}
	derived.Greet() // 输出: Hello from Derived

	// 如果想调用被遮蔽的Base的Greet方法，需要通过嵌套字段访问
	derived.Base.Greet() // 输出: Hello from Base
}
