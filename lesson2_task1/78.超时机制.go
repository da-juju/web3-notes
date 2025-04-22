package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个父context
	ctx := context.Background()

	// 创建一个带有超时的子context，超时时间为2秒
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // 取消操作，释放资源

	// 执行一个可能会长时间运行的操作
	go doSomething(ctx)

	// 等待一段时间，观察超时效果
	time.Sleep(5 * time.Second)

	// 检查context是否已经超时
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("操作超时")
	} else {
		fmt.Println("操作完成")
	}
}

func doSomething(ctx context.Context) {
	// 模拟一个长时间运行的操作
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("操作完成")
	case <-ctx.Done():
		// context超时或被取消
		fmt.Println("操作被取消或超时")
	}
}
