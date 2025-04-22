package main

import "fmt"

/*
一、概念

	数组：指一系列相同类型的数据的集合

二、数组定义

	var 数组名 [元素数量] 类型
	var arr [2] int

三、数组初始化

 1. 全部初始化
 2. 部分初始化
 3. 指定某个元素初始化
 4. 通过初始化确定数组长度
*/
func main() {

	// 数组定义: var 数组名 [元素数量]类型
	// var numberArr [5]int

	fmt.Println("//  1. 全部初始化")
	// 1. 全部初始化
	var numberArr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("numberArr[3] => %d\n", numberArr[3])

	fmt.Println("//  2. 部分初始化")
	//  2. 部分初始化
	// 使用短变量的方式，进行数组的部分初始化
	numberArr2 := [5]int{1, 2}
	fmt.Printf("numberArr2[0] => %d\n", numberArr2[0]) // 1
	fmt.Printf("numberArr2[1] => %d\n", numberArr2[1]) // 2
	fmt.Printf("numberArr2[2] => %d\n", numberArr2[2]) // 0   // 部分初始化，未赋值的索引元素为：0
	fmt.Printf("numberArr2[3] => %d\n", numberArr2[3]) // 0   // 部分初始化，未赋值的索引元素为：0
	fmt.Printf("numberArr2[4] => %d\n", numberArr2[4]) // 0   // 部分初始化，未赋值的索引元素为：0

	fmt.Println("// 3. 指定某个元素初始化")
	//  3. 指定某个元素初始化
	numberArr3 := [5]int{0: 1, 2: 3}
	fmt.Printf("numberArr3[0] => %d\n", numberArr3[0]) // 1
	fmt.Printf("numberArr3[1] => %d\n", numberArr3[1]) // 0   // 指定某个元素初始化，未赋值的索引元素为：0
	fmt.Printf("numberArr3[2] => %d\n", numberArr3[2]) // 3
	fmt.Printf("numberArr3[3] => %d\n", numberArr3[3]) // 0   // 指定某个元素初始化，未赋值的索引元素为：0
	fmt.Printf("numberArr3[4] => %d\n", numberArr3[4]) // 0   // 指定某个元素初始化，未赋值的索引元素为：0

	fmt.Println("// 4. 通过初始化确定数组长度（根据元素个数，来确定数组的长度）")
	// 4. 通过初始化确定数组长度 (根据元素个数，来确定数组的长度)
	numberArr4 := [...]int{1, 2, 4, 5}
	fmt.Printf("numberArr4[0] => %d\n", numberArr4[0]) // 1
	fmt.Printf("numberArr4[1] => %d\n", numberArr4[1]) // 2
	fmt.Printf("numberArr4[2] => %d\n", numberArr4[2]) // 4
	fmt.Printf("numberArr4[3] => %d\n", numberArr4[3]) // 5
}
