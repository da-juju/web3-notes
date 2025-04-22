package main

import "fmt"

func main() {
	//一、整数数字
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	//二、浮点数字
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)

	//三、布尔型
	var featureFlag bool = true
	fmt.Println(featureFlag)

	//四、字符串
	var firstName string = "John"
	lastName := "Doe"
	fmt.Println(firstName, lastName)

	//五、常见转义字符
	// \n：新行
	// \r：回车符
	// \t：选项卡
	// '：单引号
	// "：双引号
	// \：反斜杠
	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)

}
