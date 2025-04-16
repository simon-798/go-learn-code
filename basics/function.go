package basics

import "fmt"

type functionA struct {
	i int
}

func (a *functionA) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var function1 func(int) int

// 声明闭包
var squart2 = func(p int) int {
	p *= p
	return p
}

// Counter是一个函数，它返回另一个函数（闭包）
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func mainFunction() {

	c := Counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 2
	fmt.Println(Counter())

	for i := 0; i < 3; i++ {
		go func(j int) {
			fmt.Println(j) // 可能输出 3, 3, 3（i 是共享变量）
		}(i)
	}

	/*a := functionA{1}
	// 把方法赋值给函数变量
	function1 = a.add

	// 声明一个闭包并直接执行
	// 此闭包返回值是另外一个闭包（带参闭包）
	returnFunc := func(int) func(int, string) (int, string) {
		fmt.Println("this is a anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}(1)

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "test")
	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

	fmt.Println("a.i = ", a.i)
	fmt.Println("after call function1, a.i = ", function1(1))
	fmt.Println("a.i = ", a.i)*/
}

// 函数
func testFunction(a string, b int) (string, int) {
	return a, b
}

type functionStruct struct {
	a string
	b int
}

// 方法
func (f *functionStruct) functionMethod(a int, b string) int {
	a = 1
	b = "1"
	return a
}
