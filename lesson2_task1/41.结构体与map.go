package main

import "fmt"

type Student2 struct {
	id   int
	name string
	age  int
	addr string
}

/*
结构体与map整合
*/
func main() {

	// 1.申明结构体map
	var studentMap = make(map[int]Student2)

	// 2.结构体map初始化
	studentMap[1] = Student2{id: 1, name: "张三", age: 18, addr: "滨湖区"}
	studentMap[2] = Student2{id: 2, name: "李四", age: 30, addr: "新吴区"}
	studentMap[3] = Student2{id: 3, name: "老王", age: 44, addr: "锡山区"}

	// 3.循环遍历
	for key, val := range studentMap {
		fmt.Printf("%d=>%v\n", key, val)
	}

	fmt.Println()
	// 4.删除结构体map数据
	delete(studentMap, 2)
	for key, val := range studentMap {
		fmt.Printf("%d=>%v\n", key, val)
	}
}
