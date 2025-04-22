package main

import (
	"fmt"
	"os"
)

/*
《文件创建》

	1.创建文件
		os.Create()
	2.文件读写操作

	3.文件关闭
		defer file.Close()
*/
func main() {
	// 1.创建文件
	// a.文件路径 “D:\com\test\”
	// b.文件名 “a.txt”
	file, err := os.Create("D:\\com\\test\\a.txt")
	if err != nil {
		fmt.Println("err=", err)
	}

	// 2.文件操作
	// ......

	// 3.文件关闭(延迟调用文件关闭)
	defer file.Close()
}
