package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goRoutineTestMethod3()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goRoutineTestMethod1() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()

	go func(string) {
	}("gorouine: closure params")

	go say("in goroutine: world")

	say("hello")
}

func goRoutineTestMethod2() {
	counter := SafeCounter{}

	// 启动100个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())
}

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

func goRoutineTestMethod3() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}
