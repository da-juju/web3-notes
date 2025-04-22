package main

import "fmt"

/*
一、语法格式：

	for init; condition; post {
		// 循环体
		....
	}
		* init： 一般为赋值表达式，给控制变量赋初值；
		* condition： 关系表达式或逻辑表达式，循环控制条件；
		* post： 一般为赋值表达式，给控制变量增量或减量。
*/
func main() {

	// 循环10次
	for i := 1; i <= 10; i++ {
		fmt.Printf("循环第%d次\n", i)
	}

	// 无限循环（死循环）
	//for {
	//	fmt.Println("无线循环。。。")
	//}

	//    以下是练习案例    //
	// 练习案例（1）：1~100累加求和
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1~100累加求和：", sum)

	var evenNumberSum int
	// 练习案例（2）：1~100的偶数累加求和
	for i := 2; i <= 100; i += 2 {
		evenNumberSum += i
	}
	fmt.Println("1~100的偶数累加求和：", evenNumberSum)
}
