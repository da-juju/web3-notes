package main

import "fmt"

/*
切片传参
 1. 切片传参可以修改切片中的元素（前提是切片没有因为扩容而重新分配底层数组）
    2.可以通过append()函数对切片进行动态扩容（因为扩容而重新分配底层数组）
*/
func main() {

	var slic = make([]int, 5)
	updateSlice(slic)
	fmt.Println("元素修改slic=", slic)

	updateButSliceAppend(slic)
	fmt.Println("动态扩容len(slic)=", len(slic))
	fmt.Println("动态扩容cap(slic)=", cap(slic))
	fmt.Println("动态扩容slic=", slic)
}

/*
修改切片的第一个元素值
*/
func updateSlice(s []int) {
	s[0] = 100

}

/*
切片扩容，同时修改了切片元素，是否修改原切片元素（todo 原切片元素不变）
todo 原因：切片因扩容而重新分配底层数组（和原切片不共享内存空间）
*/
func updateButSliceAppend(s []int) {
	s = append(s, 66)
	fmt.Println("动态扩容后的len(slic)=", len(s))
	fmt.Println("动态扩容后的cap(slic)=", cap(s))
	fmt.Println("动态扩容后的slic=", s)
}
