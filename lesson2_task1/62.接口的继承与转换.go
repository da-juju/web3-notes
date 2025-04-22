package main

import "fmt"

/*
《接口继承》
	接口“继承”是通过嵌入（embedding）来实现的


	在下面代码案例中，ReadWriter 接口通过嵌入 Reader 和 Writer 接口，继承了它们的方法。
	任何实现了ReadWriter接口的类，都必须实现Read和Write方法。

*/

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

// ReadWriter 通过嵌入 Reader 和 Writer，ReadWriter 继承了它们的方法
type ReadWriter interface {
	Reader // 嵌入 Reader 接口
	Writer // 嵌入 Writer 接口
}

type Disk struct {
}

func (d *Disk) Read() {
	fmt.Println("从硬盘中读取数据")
}

func (d *Disk) Write() {
	fmt.Println("往硬盘中写入数据")
}

func main() {
	var disk Disk
	var rw ReadWriter
	rw = &disk
	rw.Read()
	rw.Write()
}
