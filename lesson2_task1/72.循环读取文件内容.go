package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
《循环读取文件内容》

	// 方法一：使用 ioutil.ReadFile 读取整个文件内容



	// 方法二：使用 os.Open 和 bufio.Scanner 逐行读取文件内容
	1.打开文件
	2.循环读取文件内容
		scanner := bufio.NewScanner(file)
		fmt.Println("\n文件内容（bufio.Scanner）:")
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

	3.关闭文件
*/
func main() {

	// 方法一：使用 ioutil.ReadFile 读取整个文件内容
	data, err := ioutil.ReadFile("D:\\com\\test\\file.txt")
	if err != nil {
		log.Fatalf("读取文件失败: %v", err)
	}
	fmt.Println("文件内容（ioutil.ReadFile）:")
	fmt.Println(string(data))

	// 1.打开文件
	file, _ := os.Open("D:\\com\\test\\file.txt")

	// 2.循环读取文件内容
	scanner := bufio.NewScanner(file)
	fmt.Println("\n文件内容（bufio.Scanner）:")
	for scanner.Scan() { // 等价于：while(scanner.Scan()){ ... }
		fmt.Println(scanner.Text())
	}

	// 3.关闭文件
	defer file.Close()
}
