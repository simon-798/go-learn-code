package lesson2

import (
	"fmt"
	"sync"
)

/*
*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func ChannelTask1() {

	num := 10
	var ch = make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 1; i <= num; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Printf("Received:%d ", v)
		}
		defer wg.Done()
	}()

	wg.Wait()

}

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

	考察点 ：通道的缓冲机制。
*/
func ChannelTask2() {
	num := 100
	var ch = make(chan int, num)
	wg := sync.WaitGroup{}

	go func() {
		for i := 1; i <= num; i++ {
			ch <- i // 发送数据到缓冲通道
		}
		close(ch) // 关键！发送完毕后必须关闭通道
	}()

	// 消费者协程：持续接收直到通道关闭
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Processed:", v)
		}
	}()

	wg.Wait()
}
