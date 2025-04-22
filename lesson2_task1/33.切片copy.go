package main

import "fmt"

/*
copy(s1,s2)

	将切片s2元素拷贝给切片s1
*/
func main() {

	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}

	// 切片拷贝：将切片s2元素拷贝给切片s1
	copy(s1, s2)
	fmt.Println("s1=", s1) // [3,4]
}
