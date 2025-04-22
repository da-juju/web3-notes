package main

import "fmt"

/*
	   切换append函数的使用
			1.append函数基本使用
			2.切片扩容
*/
func main() {

	// 创建并初始化切片
	s1 := make([]int, 4, 8) // s1= [0 0 0 0]
	fmt.Println("make([]int, 4, 8)函数创建s1=", s1)
	fmt.Println("len(s1)=", len(s1)) // 4
	fmt.Println("cap(s1)=", cap(s1)) // 8

	//1.append函数基本使用
	s1 = append(s1, 1) // s1 = [0,0,0,0,1]
	fmt.Println("append()函数追加元素后的s1=", s1)
	fmt.Println("len(s1)=", len(s1)) // 5
	fmt.Println("cap(s1)=", cap(s1)) // 8

	//2.切片扩容

	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	s1 = append(s1, 5) // s1 = [0,0,0,0,1,2,3,4,5]，s1切片cap=8，继续追加数据，s1切片会发生扩容（原cap的2倍）
	fmt.Println("扩容s1=", s1)
	fmt.Println("len(s1)=", len(s1)) // 9
	fmt.Println("cap(s1)=", cap(s1)) // 16

}
