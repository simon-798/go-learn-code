package basics

import (
	"fmt"
)

// 全局变量
var s1 = "Hello"
var zero int
var b1 = true

var (
	i  = 123
	b2 bool
	s2 = "test"
)

var (
	group = 2
)

/*
*
数据类型和变量声明
*/
func mainDataType() {
	var s = "Hello, world!"
	var bytes = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)

	var bytes2 = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var s2 = string(bytes2)
	fmt.Println(s2)

	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i        //类型推导，默认为complex128
	c3 := complex(1.10, 0.1) // c2与c3是等价的
	fmt.Println(c1 == complex64(c2))
	fmt.Println(c2 == c3)

	x := real(c2) //获取复数的实部
	fmt.Println(x)
	y := imag(c2) //获取复数的虚部
	fmt.Println(y)

	var str = string("Hello")
	fmt.Println(len(str))

	var str2 = string("Hello 世界")
	fmt.Println(len(str2))
	//中文字符，占3个字节
	var str3 = string("Hello世界")
	fmt.Println(len(str3))

	fmt.Println(group)

	name := "simon"
	fmt.Println(name)

	method1()
	method2()
	method3()
	method4()

	fmt.Println(method2())
	fmt.Println(method3())
	fmt.Println(method4())

	method()

}

func method1() {
	// 方式1，类型推导，用得最多
	a := 1
	// 方式2，完整的变量声明写法
	var b = 2
	// 方式3，仅声明变量，但是不赋值，
	var c int
	fmt.Println(a, b, c)
}

// 方式4，直接在返回值中声明
func method2() (a int, b string) {
	// 这种方式必须声明return关键字
	// 并且同样不需要使用，并且也不用必须给这种变量赋值
	return 1, "test"
}

func method3() (a int, b string) {
	a = 1
	b = "test"
	return
}

func method4() (a int, b string) {
	return
}

var a, b, c = 1, 2, 3

var e, f, g int

var h, z, j = 1, 2, "test" //类型推导

func method() {
	var k, l, m = 1, 2, 3
	var n, o, p int
	q, r, s := 1, 2, "test"
	fmt.Println(k, l, m, n, o, p, q, r, s)
}
