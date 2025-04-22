package main

import "fmt"

/*
《接口》

	在 Go 语言中，接口（interface）是一种非常重要的类型，它定义了一组方法签名，但不实现这些方法。
	接口由类型来实现，只要一个类型提供了接口中声明的所有方法，它就隐式地实现了该接口，而无需显式声明

一、接口定义

		接口通过 type 关键字和 interface 标识符来定义，其中包含了零个或多个方法的签名。
		方法签名只包括方法名、参数列表和返回类型，不包括方法体。
			type Shape interface {
	    		Area() float64
				Perimeter() float64
			}
		在这个例子中，Shape 接口定义了两个方法：Area 和 Perimeter，它们分别返回浮点数类型的结果。

二、接口实现

	在Go中，一个类型“实现”了一个接口，如果它提供了该接口中声明的所有方法的实现。不需要显式地使用 implements 关键字。

	在下面的案例中:
		我们定义了一个Nation（国籍）的接口，接口中声明了一个speak()模板方法。
		然后，我们定义了两个结构体China和America，并为它们各自实现了speak()方法.
		因此，China 和 America 类型都隐式地实现了 Nation 接口。

*/

// Nation 国籍
type Nation interface {
	speak()
}

// China 中国类
type China struct {
}

// America 美国类
type America struct {
}

/*
China实现了Nation接口
*/
func (c China) speak() {
	fmt.Println("中国人说汉语")
}

/*
America也实现了Nation接口
*/
func (a2 America) speak() {
	fmt.Println("美国人说英语")
}

func main() {
	china := China{}
	china.speak()

	usa := America{}
	usa.speak()
}
