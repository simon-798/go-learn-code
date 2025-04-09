package main

import (
	"fmt"
	//"time"
)

/*func main() {
	var a int
	if b := 1; b == 0 {
		fmt.Println("b == 0")
	} else {
		c := 2
		fmt.Println("declare c = ", c)
		fmt.Println("b == 1")
	}
	// fmt.Println(b)
	// fmt.Println(c)

	switch d := 3; d {
	case 1:
		e := 4
		fmt.Println("declare e = ", e)
		fmt.Println("d == 1")
	case 3:
		f := 4
		fmt.Println("declare f = ", f)
		fmt.Println("d == 3")
	}
	// fmt.Println(e)
	// fmt.Println(f)

	for i := 0; i < 1; i++ {
		forA := 1
		fmt.Println("forA = ", forA)
	}
	// fmt.Println("forA = ", forA)

	select {
	case <-time.After(time.Second):
		selectA := 1
		fmt.Println("selectA = ", selectA)
	}
	// fmt.Println("selectA = ", selectA)

	// 匿名代码块
	{
		blockA := 1
		fmt.Println("blockA = ", blockA)
	}
	// fmt.Println("blockA = ", blockA)

	fmt.Println("a = ", a)
}*/

/*var a int

func main() {
	{
		fmt.Println("global variable, a = ", a)
		a = 3
		fmt.Println("global variable, a = ", a)

		a := 10
		fmt.Println("local variable, a = ", a)
		a--
		fmt.Println("local variable, a = ", a)
	}
	fmt.Println("global variable, a = ", a)
}*/

func main() {
	var b int = 4
	fmt.Println("local variable, b = ", b)
	if b := 3; b == 3 {
		fmt.Println("if statement, b = ", b)
		b--
		fmt.Println("if statement, b = ", b)
	}
	fmt.Println("local variable, b = ", b)
}
