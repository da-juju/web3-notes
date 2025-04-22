package main

import (
	"fmt"
	"os"
)

/*
《文件写入数据》

	1.WriteString：写入字符串
	2.Write：写入字节切片
	3.WriteAt：从指定位置开始写入数据
	4. file.Seek(0,io.SeekEnd)   // 获取文件中写入数据的长度
*/
func main() {
	WriteStr()
	Write()
	WriteAt()
}

func WriteStr() {
	file, err := os.Create("D:\\com\\test\\b.txt")
	if err != nil {
		fmt.Println("err=", err)
	}
	file.WriteString("Hello World")
	defer file.Close()
}

func Write() {
	file, err := os.Create("D:\\com\\test\\c.txt")
	if err != nil {
		fmt.Println("err=", err)
	}
	var str = "Hello World2"
	file.Write([]byte(str))
	defer file.Close()
}

func WriteAt() {
	file, err := os.Create("D:\\com\\test\\d.txt")
	if err != nil {
		fmt.Println("err=", err)
	}
	file.WriteString("Hello World")
	file.WriteAt([]byte("aaa"), 11)
	defer file.Close()
}
