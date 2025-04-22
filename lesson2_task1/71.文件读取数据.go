package main

import (
	"fmt"
	"os"
)

/*
《文件读取器数据》

	1.打开要读取的文件
		func Open(name string) (*File, error) {
			// 	os.O_APPEND（追加）、0：没有任何权限
			return OpenFile(name, O_RDONLY, 0)
		}
	2.文件读取
	3.文件关闭
*/
func main() {
	//1.打开要读取的文件
	file, err := os.Open("D:\\com\\test\\a.txt")
	if err != nil {
		fmt.Println("err=", err)
	}

	//2.文件读取
	var buffer = make([]byte, 2*1024)           // 存储数据的字节切片
	n, _ := file.Read(buffer)                   // 将文件数据读取到buffer中
	fmt.Println("文件读取的数据：", string(buffer[:n])) // 切片截取有效数据

	//3.文件关闭
	defer file.Close()
}
