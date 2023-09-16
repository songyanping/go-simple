package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	// 模拟工作
	time.Sleep(time.Second)
	result := id * 2

	// 将结果发送到通道
	resultChan <- result
}

func main() {
	numWorkers := 5 // 假设有5个工作goroutine

	// 创建等待组和结果通道
	var wg sync.WaitGroup
	resultChan := make(chan int, numWorkers)

	// 启动工作goroutine
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, resultChan)
		
	}

	// 等待所有工作goroutine完成
	wg.Wait()

	// 关闭结果通道（一定要在所有工作goroutine都完成后关闭通道）
	close(resultChan)

	// 从通道中收集并打印结果
	for result := range resultChan {
		fmt.Printf("Worker result: %d\n", result)
	}
}

//func main() {
//	var wg sync.WaitGroup
//	// 开 N 个后台打印线程
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//
//		go func() {
//			fmt.Printf("你好, 世界 %d\n", i)
//			wg.Done()
//		}()
//	}
//	// 等待 N 个后台线程完成
//
//	wg.Wait()
//}
