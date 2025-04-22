package main

import (
	"fmt"
	"strings"
)

/*
《字符串常用方法》

	1.Contains()：字符串中是否包含
	2.Join()：字符串连接
	3.Index()：获取某个字符串的下标位置（不存在index返回：-1）
	4.Repeat()：字符串重复多少次
	5.Replace()：字符串替换{Replace(str, "a", "A", 3)=>需要替换的字符串，旧值，新值，需要替换的次数（-1：全部替换）}
	6.Split()：字符串分割（返回字符串切片）
*/
func main() {

	var str = "sadasd1sadsa"
	//1.Contains()
	isContains := strings.Contains(str, "a")
	fmt.Println("isContains=", isContains)

	//2.Join()
	var strArr = []string{"4_1", "4_2"}
	strJoin := strings.Join(strArr, ",")
	fmt.Println("strJoin=", strJoin)

	//3.Index()
	index := strings.Index(str, "12")
	fmt.Println("index=", index) // 字符串的索引位置（不存在index返回：-1）

	//4.Repeat()
	repeat := strings.Repeat("go", 3)
	fmt.Println("repeat=", repeat) // gogogo

	//5.Replace()
	replace := strings.Replace(str, "a", "A", -1)
	fmt.Println("replace=", replace)

	//6.Split()
	splitAfter := strings.SplitAfter(str, "")
	fmt.Println("splitAfter=", splitAfter)
}
