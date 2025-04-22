package main

import "fmt"

/*
切片指针

	切片指针，指向切片的指针。这允许你通过指针传递切片，从而允许函数修改原始切片的内容。
	1.声明与初始化切片指针
		使用 * 操作符来声明一个指向切片的指针。例如，*[]int 是一个指向 []int 类型的数组的指针。
			①申明：
				var slicePtr *[]int
			②初始化：
				slice := []int{1, 2, 3}
				slicePtr = &slice
	2.值的修改
	3.循环遍历
*/
func main() {

	// 1.声明与初始化切片指针
	fmt.Println("//  1.声明与初始化切片指针")
	// 申明切片指针
	var slicePtr *[]int
	// 初始化
	slice := []int{1, 2, 3}
	slicePtr = &slice

	// 2.值的获取和修改
	fmt.Println("//  2.值的获取和修改")
	fmt.Println("获取切片所有元素：", *slicePtr) // [1 2 3]
	fmt.Println("获取第一个元素：", (*slicePtr)[0])
	// 修改第一个元素
	(*slicePtr)[0] = 100
	fmt.Println("修改第一个元素：", *slicePtr) // [100 2 3]

	// 3.循环遍历
	for index, value := range *slicePtr {
		fmt.Printf("%d=>%d\n", index, value)
	}
}
