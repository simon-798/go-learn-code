package main

import "fmt"

/*
*
流程控制
*/
/*func main() {

	var a int = 10
	if b := 1; a > 10 {
		b = 2
		// c = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("其他")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println(b)
		fmt.Println(c)
	}

}*/

type CustomType struct {
	name string
}

func main() {
	a := "test string"

	// 1. 基本用法
	switch a {
	case "test":
		fmt.Println("a = ", a)
	case "s":
		fmt.Println("a = ", a)
	case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
		fmt.Println("catch in a test, a = ", a)
	case "n":
		fmt.Println("a = not")
	default:
		fmt.Println("default case")
	}

	// 变量b仅在当前switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3, 4:
		fmt.Println("b = 3 or 4")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b = ", b)
	}

	// 不指定判断变量，直接在case中添加判定条件
	b := 5
	switch {
	case a == "t":
		fmt.Println("a = t")
	case b == 3:
		fmt.Println("b = 5")
	case b == 5, a == "test string":
		fmt.Println("a = test string; or b = 5")
	default:
		fmt.Println("default case")
	}

	var d interface{} //等价于java中的Object类型
	//var e int = 1
	d = 1

	/*cc := CustomType{name: "simon"}
	d = &cc*/

	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case *byte:
		fmt.Println("d is byte point type, ", t)
	case *int:
		fmt.Println("d is int type, ", t)
	case *string:
		fmt.Println("d is string type, ", t)
	case *CustomType:
		fmt.Println("d is CustomType pointer type, ", t)
	case CustomType:
		fmt.Println("d is CustomType type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}
}
