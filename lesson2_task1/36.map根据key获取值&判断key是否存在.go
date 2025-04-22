package main

import "fmt"

/*
根据key获取value，并判断key是否存在

	格式：val, isExist := map[key]
		1. val：key对应的value
		2. iExist：判断key是否存在
*/
func main() {

	m1 := make(map[string]int)

	m1["a"] = 1
	m1["b"] = 2
	m1["c"] = 3

	value := m1["c"]
	fmt.Println(value)
	val, isExist := m1["c"]
	fmt.Printf("%d=>%t\n", val, isExist)

	val2, isExist2 := m1["d"]
	fmt.Printf("%d=>%t\n", val2, isExist2)
}
