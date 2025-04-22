package main

import "fmt"

/*
数组遍历

		1.for...len()
	    2.for...range
*/
func main() {
	// 定义数组
	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(" // 1.for...len()遍历数组 ")
	// 1.for...len()
	for i := 0; i < len(arr); i++ {
		fmt.Printf("下标：%d ==> 元素：%d\n", i, arr[i])
	}

	fmt.Println(" // 2.for...range遍历数组 ")
	for idx, val := range arr {
		fmt.Printf("下标：%d ==> 元素：%d\n", idx, val)
	}
}
