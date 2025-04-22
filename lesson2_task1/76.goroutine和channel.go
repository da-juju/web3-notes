package main

import "fmt"

func sum1(a []int, ch chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	ch <- total // 将结果发送到 Channel
}

// "%"和"/"的用法
// "%"是求余数，"/"是求商
// "%"是求余数示例：10 % 3 = 1
// "/"是求商示例：10 / 3 = 3
//func main() {
//	a := []int{1, 2, 3, 4, 5}
//	ch := make(chan int)
//	go sum1(a[:len(a)/2], ch) // 从数组的索引 0 开始，到索引 2 结束，不包括 2
//	go sum1(a[len(a)/2:], ch) // 从数组的索引 2 开始，到索引 5 结束，包括 5
//	x, y := <-ch, <-ch        // 主goroutine从channel ch中接收两个值，分别赋给x和y
//	fmt.Println(x, y, x+y)
//}

//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	go func() {
//		ch1 <- 1
//	}()
//
//	go func() {
//		ch2 <- 2
//	}()
//
//	select {
//	case msg1 := <-ch1:
//		fmt.Println("Received from ch1:", msg1)
//	case msg2 := <-ch2:
//		fmt.Println("Received from ch2:", msg2)
//	}
//}

//func main() {
//
//	ch := make(chan int)
//
//	// 子goroutine向通道中发送数据
//	go func() {
//		for i := 0; i < 5; i++ {
//			ch <- i
//		}
//
//		// 关闭通道
//		close(ch)
//	}()
//
//	for {
//		// 从通道中接收数据
//		if v, ok := <-ch; ok {
//			fmt.Println(v)
//		} else {
//			break
//		}
//	}
//	fmt.Println("main over")
//}

// channel for...range遍历
//func f1(ch chan<- int) {
//	fmt.Println("f1.start")
//	for i := 0; i < 100; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
//func f2(ch1 <-chan int, ch2 chan<- int) {
//	fmt.Println("f2.start")
//	for val := range ch1 {
//		ch2 <- val * 2
//	}
//	close(ch2)
//}
//
//func main() {
//	ch1 := make(chan int, 100)
//	ch2 := make(chan int, 100)
//
//	go f1(ch1)
//	go f2(ch1, ch2)
//
//	for val := range ch2 {
//		fmt.Println(val)
//	}
//}

//channel select的多路复用

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // 读取
			fmt.Println(x)
		case ch <- i: // 写入
		}
	}
}
