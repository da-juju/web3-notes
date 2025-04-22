package main

import "fmt"

/*
map中，根据key删除value

	// map：map集合的变量名称
	// key：要删除的key
	delete(map, key)
*/
func main() {
	m1 := map[string]int{"zs": 1, "ls": 2, "ww": 3}
	fmt.Println("m1=", m1)

	// 根据key删除value
	delete(m1, "zs")
	fmt.Println("删除key=zs的m1=", m1)
}
