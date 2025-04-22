package main

import "fmt"

/*
《”类“的继承》

一、概念

	从一个类派生另一个类，继承其属性和方法。可以通过匿名字段（anonymous field）来实现类似的效果。

二、类的继承创建对象的初始化

	student1 := Student{
			Person{
				id:   1,
				name: "张三",
				age:  18,
			},
			88.5,
	}
*/
func main() {

	// 一、通过匿名字段来实现类的继承
	// 人员类
	type Person struct {
		id   int
		name string
		age  int
	}

	// 学生类继承人员类
	type Student struct {
		Person         // 匿名字段，表示嵌入 Person 结构体
		grade  float64 // 学生成绩
	}

	// 教师类继承人员类
	type Teacher struct {
		Person         // 匿名字段，表示嵌入 Person 结构体
		subject string // 教学科目
	}

	// 二、类的继承创建对象的初始化
	// 1.全部初始化
	student1 := Student{Person{1, "张三", 18}, 88.5}
	fmt.Println("1.全部初始化student1=", student1)
	// 2.部分初始化
	// 2.1.部分初始化子类属性
	student2 := Student{grade: 90}
	fmt.Println("2.1.部分初始化子类属性student2=", student2)
	// 2.2.部分初始化父类属性
	student3 := Student{Person: Person{id: 3, name: "苍老师"}}
	fmt.Println("2.2.部分初始化父类属性student3=", student3)
}
