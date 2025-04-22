package main

import "fmt"

/*
一、二维数组概念

	二维数组：是一个二维的表格结构
		如：var arr2 [2][3]int
		表示一个2行3列的表格，
	           0      1     2
			————————————————————
		0	|  123 |     |     |
			————————————————————
		1	|      |     | 45  |
			————————————————————
		arr2[0][0] = 123;
		arr2[1][2] = 45;

二、二维数组初始化

	1.全部初始化
	2.部分初始化
	3.指定元素初始化
	4.通过初始化确定二维数组行数

三、二维数组遍历
*/
func main() {

	// 1.全部初始化
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr1 = ", arr1)

	// 2.部分初始化
	var arr2 [2][3]int = [2][3]int{{1, 2}, {5, 6}}
	fmt.Println("arr2 = ", arr2)

	// 3.指定元素初始化
	var arr3 [2][3]int = [2][3]int{0: {0: 1, 2: 3}} // arr3[0][0] = 1; arr3[0][2] = 3
	fmt.Println("arr3 = ", arr3)

	// 4.通过初始化确定二维数组行数
	var arr4 [2][3]int = [...][3]int{{}, {4, 5, 6}} // arr4[1][0] = 4; arr4[1][1] = 5; arr4[1][2] = 6
	fmt.Println("arr4 = ", arr4)

	//  遍历二维数组    //
	fmt.Println("//    遍历二维数组    //")
	for index, arrEle := range arr4 {
		for index2, ele2 := range arrEle {
			fmt.Printf("arr4[%d][%d] = %d\n", index, index2, ele2)
		}
	}
}
