package basics

import (
	"fmt"
	"time"
)

func mainSelectTest() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("ch: %d\n", i)
			ch1 <- i
			ch2 <- i
			ch3 <- i

		}
	}()

	timeout := time.After(2 * time.Second)
	for {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		case z := <-ch3:
			fmt.Printf("receive %d from channel 3\n", z)
		case <-timeout:
			fmt.Println("操作超时")
			return
		}
	}
}
