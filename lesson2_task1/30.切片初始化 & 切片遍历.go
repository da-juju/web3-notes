package main

import "fmt"

/*
一、切片概念

	在 Go 语言（Golang）中，切片（slice）是一种动态数组，其长度可以根据需要增长和缩小。
	与数组不同，切片不需要在声明时指定其大小，且其大小可以在后续操作中改变。切片是对底层数组的抽象和封装，提供了更灵活和方便的内存管理方式。

二、切片创建

			1.使用var
				# 默认是空的切片
		        # 长度为0
			2.使用类型推导（短变量）的方式
				# 默认是空的切片
	        	# 长度为0
			3.使用make()函数
				# 基本语法：make(切片类型，长度，容量)
				# make()函数创建切片的特点：
					① 长度指已经初始化的空间，容量指已经开辟了内存的空间（但未使用的空闲空间）
					② 长度不可大于容量
					③ len(sliceVar)：获取长度、cap(sliceVar)：获取容量
					④ make()函数可以省略容量参数，默认：容量=长度

三、切片初始化

	1.使用var
	2.使用类型推导（短变量）的方式
	3.使用make()函数
*/
func main() {

	// 二、切片创建[ 1.var、2.短变量、3.make(类型,长度,容量) ]
	fmt.Println("//  二、切片创建   //")
	//1.使用var
	var slice1 []int
	fmt.Println(slice1)

	//2.使用类型推导（短变量）的方式
	slice2 := []int{}
	fmt.Println(slice2)

	//3.使用make()函数
	slice3 := make([]int, 2, 5)
	fmt.Println(slice3)

	// 三、切片初始化
	fmt.Println("//  三、切片初始化   //")
	fmt.Println("//  使用append()函数对切片进行初始化")
	slice2 = append(slice2, 1, 2, 3, 4, 5)
	fmt.Println("对切片slice2初始化：", slice2)

	fmt.Println("//  使用append()函数对make()函数方式创建的切片进行初始化")
	// 对slice3切片下标为0的元素进行赋值
	slice3[0] = 10

	// 使用append()进行值追加
	slice3 = append(slice3, 55)
	fmt.Println("对切片slice3追加元素：", slice3)

	// 切片遍历
	fmt.Println("//  切片遍历")
	for index, ele := range slice3 {
		fmt.Printf("%d => %d\n", index, ele)
	}
}
