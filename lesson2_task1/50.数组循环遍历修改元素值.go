package main

import "fmt"

/*
for...i 和for...range 两种方式循环遍历修改数组中的元素值
 1. for...i：循环遍历，通过arr[i]下标方式，【可修改】数组元素值
 2. for...range：循环遍历，通过修改range得到的value，【不可以修改】数组元素值
*/
func main() {

	//  1. for...i：循环遍历，通过arr[i]下标方式，【可修改】数组元素值
	var arr = [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		arr[i] = 10
	}
	fmt.Println("arr=", arr)

	// 2. for...range：循环遍历，通过修改range得到的value，【不可以修改】数组元素值
	var arr2 = [3]int{1, 2, 3}
	for _, value := range arr2 {
		value = 10
		fmt.Println("value=", value)
	}
	fmt.Println("arr2=", arr2)

}
