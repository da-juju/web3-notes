package main

import "fmt"

/*
《面向对象》

	Go语言虽然并没有像传统面向对象编程语言（如Java、C++）那样直接提供“类”（class）、“继承”（inheritance）等概念。
	但它确实支持面向对象编程的核心理念，并通过其独特的语法结构实现了OOP的三大特性：封装、继承、多态

一、面向对象编程的基本概念

	对象：具有状态和行为的实体。在Go中，对象通常通过结构体（struct）来表示。
	类：定义对象蓝图的模板。在Go中，虽然没有直接的“类”概念，但结构体（struct）可以看作是一个类的替代品，它定义了对象的属性和方法。
	方法：作用于对象的函数。在Go中，方法是通过在函数定义前添加接收者（receiver）来实现的，接收者可以是值接收者或指针接收者。
	继承：从一个类派生另一个类，继承其属性和方法。可以通过组合（embedding）或匿名字段（anonymous field）来实现类似的效果。
	多态：根据对象类型调用不同实现的相同方法。在Go中，多态是通过接口（interface）来实现的。

二、Go语言中的面向对象特性

	1.结构体（struct）：
		用于定义类型的语法结构。
		可以看作是一个类的替代，它包含了对象的属性和方法。
		例如，定义一个“书本”结构体：
			type Books struct {
				title   string
				author  string
				subject string
				book_id int
			}
	2.方法（method）：
		func (b Books) sale() {
			fmt.Println(b.title, "出售一本书")
		}
	3.接口（interface）：
		指定类型必须实现的方法集合。
		在Go中，接口是实现多态的关键。任何类型只要实现了接口中定义的方法，就可以被认为实现了该接口。
		例如，定义一个Speaker接口：
			type Speaker interface {
				Speak() string
			}

	4.
*/
func main() {

	// 在Go中，虽然没有直接的“类”概念，但结构体（struct）可以看作是一个类的替代品，它定义了对象的属性和方法。
	// 以下的”Book“结构体，可以理解为java中的Class类（定义对象蓝图的模板）
	type Books struct {
		title   string
		author  string
		book_id int
	}

	// 创建对象
	book1 := Books{book_id: 1, title: "格林童话", author: "雅各布·格林和威廉·格林兄弟"}
	book2 := Books{book_id: 1, title: "老人与海", author: "欧内斯特·米勒尔·海明威"}
	fmt.Println("book1=", book1)
	fmt.Println("book2=", book2)
}
