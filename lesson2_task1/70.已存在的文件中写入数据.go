package main

import (
	"fmt"
	"os"
)

/*
《已有文件写入数据》
	func os.OpenFile(name string, flag int, perm FileMode) (*File, error)
	1.参数一：
		name string：要打开的文件名（包括路径）
	2.flag：打开文件的模式标志
		常用的标志包括
			os.O_RDONLY（只读）、
			os.O_WRONLY（只写）、
			os.O_RDWR（读写模式）、
			os.O_APPEND（追加）、
	3.perm FileMode：操作权限
		0：没有任何权限
		1：执行权限（如果是可执行文件，是可以运行的）
		2：写权限
		3：写权限与执行权限
		4：读权限
		5：读权限与执行权限
		6：读写权限
		7：读、写与执行权限
*/

func main() {
	file, err := os.OpenFile("D:\\com\\test\\a.txt", os.O_APPEND, 6)
	if err != nil {
		fmt.Println("err=", err)
	}
	defer file.Close()
	// 通过文件指针，读写操作
	file.WriteString("追加数据")
}
