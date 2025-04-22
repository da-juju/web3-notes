package main

import "fmt"

/*
1. 输入考试成绩
2. 大于等于90分，优秀
3. 大于等于80小于90，中等
4. 大于等于60小于80，一般
5. 小于60，不合格
*/
func main() {
	var score int
	fmt.Println("输入你的考试成绩：")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 && score < 90 {
		fmt.Println("中等")
	} else if score >= 60 && score < 80 {
		fmt.Println("一般")
	} else {
		fmt.Println("不合格")
	}
}
