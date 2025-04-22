package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex // 互斥锁
)

func worker() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x += 1
		lock.Unlock() // 解锁
	}
	wg.Done() // 通知main函数工作已经完成
}

func main() {
	wg.Add(2) // 通知main函数要等待2个goroutine
	go worker()
	go worker()
	wg.Wait() // 等待两个goroutine结束
	fmt.Println(x)
}
