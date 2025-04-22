package main

import "fmt"

func main() {

	// 输出1，2，4，5
	// 使用continue语句，跳出本次循环，继续下次循环
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("数值：%d\n", i)
	}

	// 换行
	fmt.Println()

	// break：直接结束本次循环
	for i := 1; ; i++ {
		fmt.Printf("输出%d次\n", i)
		if i == 10 {
			break
		}
	}
}
