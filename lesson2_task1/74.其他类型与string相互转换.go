package main

import (
	"fmt"
	"strconv"
)

/*
《其他类型与string相互转换》

	1.其他类型转string
	2.string转其他类型
*/
func main() {
	//1.其他类型转string
	//bool --> string
	formatBool := strconv.FormatBool(true)
	fmt.Println("formatBool=", formatBool)
	//int --> string
	itoa := strconv.Itoa(123)
	fmt.Println("itoa=", itoa)

	//2.string转其他类型
	// string --> bool
	parseBool, _ := strconv.ParseBool("true") // 如果转换失败，则错误不为 nil
	fmt.Println("parseBool=", parseBool)

	// string --> int
	atoi, _ := strconv.Atoi("123") //如果转换失败，则错误不为 nil
	fmt.Println("atoi=", atoi)
}
