package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2}
	// 数组作为函数的形参
	arrPrint(arr1)
}

func arrPrint(arr [5]int) {
	for idx, val := range arr {
		fmt.Printf("下标：%d ==> 元素：%d\n", idx, val)
	}
}
