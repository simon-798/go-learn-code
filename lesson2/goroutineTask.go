package lesson2

import (
	"fmt"
	"sync"
	"time"
)

/*
*
1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func RunGoroutineTask1() {

	var wg = sync.WaitGroup{}
	wg.Add(2) //加入两个协程计数器

	go func() {
		defer wg.Done() //计数器减一
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("输出1到10之间的奇数:", i)
			}
		}
	}()

	go func() {
		defer wg.Done() //计数器减一
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("输出1到10之间的偶数:", i)
			}
		}
	}()

	wg.Wait()
}

/*
*
2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func RunGoroutineTask2() {

	// 创建测试任务集合
	// 每个任务用匿名函数表示，这里用sleep模拟耗时操作
	tasks := []func(){
		func() {
			time.Sleep(1 * time.Second) // 任务1
		},
		func() {
			time.Sleep(2 * time.Second) // 任务2
		},
		func() {
			time.Sleep(3 * time.Second) // 任务3
		},
	}

	// 调用调度器执行任务
	durationArray := Schedule(tasks...)

	for i, duration := range durationArray {
		// i+1 转换为人类可读的序号（任务从1开始）
		// %v 自动选择合适的时间格式显示
		fmt.Printf("任务 %d 耗时：%v \n", i+1, duration)
	}

}

// 任务调度器函数（核心逻辑）
// 参数：可变长度的函数切片（每个函数代表一个任务）
// 返回：每个任务的耗时统计结果（按原始顺序）
func Schedule(tasks ...func()) []time.Duration {

	// 定义结果结构体，用于保存任务索引和对应耗时
	type result struct {
		index    int           // 记录任务原始顺序
		duration time.Duration // 记录任务执行耗时
	}

	// 创建带缓冲的结果通道（重要并发控制）
	// 缓冲大小等于任务数，保证所有任务都能立即发送结果
	resultInChan := make(chan result, len(tasks)) //带缓冲的通道

	// 并发执行任务的核心逻辑
	// 遍历所有任务，为每个任务启动独立协程
	for index, vFunc := range tasks {

		// 使用立即执行函数避免闭包陷阱
		// 通过参数传递当前索引和任务副本（关键细节）
		go func(index int, vFuncParam func()) {
			start := time.Now()           // 记录任务开始时间
			vFuncParam()                  // 这里执行的是传入的vFuncParam参数，最终都是在执行我们在RunGoroutineTask2函数中定义的匿名函数
			duration := time.Since(start) // 计算耗时
			resultInChan <- result{
				index:    index,    // 保留原始索引
				duration: duration, // 计算精确耗时
			}
		}(index, vFunc) // 显式传递循环变量（避免共享变量问题）
	}

	// 结果收集切片
	resultDurationArray := make([]time.Duration, len(tasks))

	timeout := time.After(5 * time.Second)
	// 循环接收所有任务结果（循环次数=任务数量）
	for range tasks {
		// 使用select语句接收结果（可扩展超时控制）
		select {
		case result := <-resultInChan: //接收chan的数据
			// 按照原始索引位置存储结果（关键设计）
			// 保证无论协程完成顺序如何，结果都能正确对应
			resultDurationArray[result.index] = result.duration
		case <-timeout:
			fmt.Println("超时处理")
			return nil
		}
	}

	return resultDurationArray // 返回有序结果
}
