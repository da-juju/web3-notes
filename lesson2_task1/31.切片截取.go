package main

import "fmt"

/*
切片（slice）截取是指从一个已有的切片中创建一个新的切片，新切片将包含原切片中的一部分或全部元素。
新切片引用与原切片相同的底层数组，但具有不同的长度和容量。（共享底层数组）

	注意：新切片和原切片可能共享同一个底层数组。todo 导致：对新切片中元素的修改，可能会影响原切片（反之亦然）

（但这仅当它们共享相同的底层数组段时成立）

			s[low:high:max]
				# low：从切片s截取的开始位置
				# high：从切片s截取的结束位置（但是不包含high下标对应的元素）
				# max：用于计算截取后的切片的容量
		           cap(s1) = max - low
	               len(s1) = high - low
*/
func main() {

	// 创建并初始化切片
	s := []int{1, 2, 3, 4, 5, 6}

	// 对s切片截取，得到新的切片s2
	s2 := s[0:3:5]
	fmt.Println(s2)

	//语法格式为：
	//// s[low:high]，其中 low 为起始索引（包含），high 为结束索引（不包含）。
	// 其他截取操作
	// s[n]： 切片s中索引位置为n的项
	// s[:]：从切片s的索引位置0到len(s)-1处所获得的切片
	// s[low:]：从切片s的索引位置low到len(s)-1处所获得的切片
	// s[:high]：从切片s的索引位置0到high处所获得的切片，len=high
	// s[low:high]：从切片s的索引位置low到high处所获得的切片，len=high-low
	// s[low:high:max]：从切片s的索引位置low到high处所获得的切片，len=high-low、cap=max-low
}
