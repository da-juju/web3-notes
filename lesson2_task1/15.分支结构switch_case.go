package main

import "fmt"

func main() {

	var days int
	fmt.Println("请输入数字1~7：")
	fmt.Scan(&days)

	switch days {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	default:
		fmt.Println("周末")
	}

}
