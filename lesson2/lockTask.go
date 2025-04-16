package lesson2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func LockTask1() {

	counter := Counter{}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter.incrementCount(&wg)
	}

	//等待所有协程执行完毕
	//time.Sleep(time.Second)
	wg.Wait()

	fmt.Println("counter is", counter.count)
}

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) incrementCount(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	for i := 0; i < 1000; i++ {
		c.count++

	}
	c.mu.Unlock()
}

/*
*	题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func LockTask2() {
	counter := atomic.Int64{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Add(1)
			}

		}()
	}

	wg.Wait() //等待所有协程执行完毕
	fmt.Println("counter is", counter.Load())
}

type UnsafeCounter struct {
	count int
}

/*
*	题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
写法二
*/
func LockTask3() {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter is", counter)
}
