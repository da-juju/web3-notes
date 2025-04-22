package main

import "fmt"

/*
一、深浅拷贝

	1.浅拷贝：
			复制对象的引用地址（指针），不赋值对象本身，
			新旧对象共享同一块内存空间。
	2.深拷贝：
			复制并创建一个一摸一样的对象，
			不共享内存，新对象修改，不影响旧对象的数据。

二、案例分析：

 1. ap := &num  // 0x1d06098
 2. p *int      // 0x1d06098
 3. ap和p内存地址，指向同一块内存空间（num=10）,
    因此通过解引用的方式修改值（*p = 69），会影响num的值
*/
func main() {

	var num = 10
	ap := &num
	fmt.Println("ap=", ap)   // 0x1d06098
	fmt.Println("&ap=", &ap) // 0x1d08078
	update(ap)
	fmt.Println(num)
}

func update(p *int) {
	fmt.Println("p=", p)   // 0x1d06098
	fmt.Println("&p=", &p) // 0x14260a8
	// 解引用修改遍历
	*p = 69
}
