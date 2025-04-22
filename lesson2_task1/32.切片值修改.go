package main

import "fmt"

/*
切片的值修改

	新切片引用与原切片相同的底层数组，但具有不同的长度和容量。（共享底层数组）
	todo 导致：对新切片中元素的修改，可能会影响原切片（反之亦然）
*/
func main() {

	// 切片创建并初始化
	ints := []int{1, 2, 3, 4, 5, 8}
	fmt.Println("原切片ints=", ints)

	// 截取
	intsSub := ints[1:3] // [2,3]
	fmt.Println("截取切片intsSub=", intsSub)

	// 截取切片的值修改
	intsSub[0] = 10
	fmt.Println("值修改后的原切片ints=", ints)        // [1,10,3,4,5,8]
	fmt.Println("值修改后的截取切片intsSub=", intsSub) // [10,3]
}
