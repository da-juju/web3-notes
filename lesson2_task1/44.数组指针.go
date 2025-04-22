package main

import "fmt"

/*
数组指针

	一、概念和基本使用

		1.概念：数组指针：指一个数组的指针（也就是指针的内存地址）
		2.数组指针申明
			var 数组指针变量 *[下标]类型
			var p *[5]int

	二、将数组指针作为函数参数
*/
func main() {

	fmt.Println("// 一、概念和基本使用  //")
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var p *[5]int
	p = &arr
	fmt.Println("数组指针：", p)
	// 通过数组指针获取数组元素
	// *p：是一个解引用操作，用于访问指针p所指向的变量的值
	fmt.Println(" 通过数组指针获取数组元素:", *p)

	// 通过数组指针p，获取数组中某个元素
	// todo 注意：[]的优先级高于p
	fmt.Println("通过数组指针p，获取数组中某个元素：", (*p)[0])
	// 数组指针获取数组中的元素"简写方式"(切片没有这种简写方式)
	fmt.Println("数组指针获取数组中的元素\"简写方式\"：", p[1])

	fmt.Println("// 二、将数组指针作为函数参数  //")
	updateByArrayPointer(p)
	fmt.Println("通过数组指针修改数组元素：", arr)

}

/*
通过数组指针，修改数组的元素
*/
func updateByArrayPointer(pt *[5]int) {
	pt[0] = 100
}
