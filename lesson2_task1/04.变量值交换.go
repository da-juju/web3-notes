package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20
	var tem int

	tem = b
	b = a
	a = tem

	// 10 20 -> 20 10
	fmt.Println(a, b)

	// 20 10 -> 10 20
	a, b = b, a
	fmt.Println(a, b)
}
