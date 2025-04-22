package main

import "fmt"

/*
统计一个英文字母的字符串，每个字母出现的次数
*/
func main() {
	str := "abdsdsddsadsa"
	// key=字母
	// value=出现次数
	strMap := make(map[rune]int)
	for _, char := range str {
		strMap[char] += 1
	}

	// 遍历map输出
	fmt.Println(strMap)
	for key, val := range strMap {
		fmt.Printf("%c=>%d\n", key, val)
	}
}
