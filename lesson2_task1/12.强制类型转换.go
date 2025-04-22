package main

import "fmt"

/*
强制类型转换
float --> int
*/
func main() {
	// float --> int
	var money float64 = 3.15
	fmt.Printf("%d", int(money))
}
