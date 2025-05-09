package main

import (
	_ "go_learn_code/basics" // 用下划线标记未使用的包，不会触发编译错误
	"go_learn_code/lesson3"
)

func main() {

	/*numArray1 := []int{1, 2, 3, 4, 5, 6, 108, 6, 5, 4, 3, 2, 1}
	result := lesson1.Job136Method(numArray1)
	fmt.Println(result)

	numArray2 := []int{2, 7, 9, 3, 1}
	ret2 := lesson1.Job198Method(numArray2)
	fmt.Println(ret2)

	lesson1.Job21Method()

	lesson1.Job46Method()

	lesson1.Job344Method2()

	lesson1.Job69Method()

	lesson1.Job26Method()

	lesson1.Job56Method()

	lesson1.Job430Method()

	lesson1.Job729Method()

	//lesson2 task
	var num = 8
	lesson2.PointerTask1(&num)
	fmt.Println("num:", num)

	nums := []int{1, 2, 3, 4, 5}
	lesson2.PointerTask2(nums)
	fmt.Println("nums:", nums)

	//在 Go 中，主函数（main）退出时会立即终止所有正在运行的 goroutine。
	//启动了两个 goroutine，但主函数没有等待它们完成，导致程序提前退出，因此会看不到输出。
	//第一种方法使用time.Sleep(1 * time.Second)，等待足够的时间，让两个协程能执行完
	//第二种方法使用同步机制：使用 sync.WaitGroup 确保主函数等待所有 goroutine 完成。
	lesson2.RunGoroutineTask1()
	time.Sleep(1 * time.Second)

	lesson2.RunGoroutineTask2()

	lesson2.InterfaceTask1()

	lesson2.InterfaceTask2()

	lesson2.ChannelTask1()

	lesson2.ChannelTask2()

	lesson2.LockTask1()

	lesson2.LockTask3()*/

	//lesson3.GormTest()

	//lesson3.Job1Method()

	//lesson3.Job2Method()

	//lesson3.Job3Method()

	//lesson3.Job4Method()

	//lesson3.Job5Method()
	//lesson3.Job6Method()
	lesson3.Job7Method()

}
