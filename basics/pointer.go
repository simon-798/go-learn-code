package main

import (
	"fmt"
	"unsafe"
)

/*
*
指针
*/
func main() {
	//pointMethod1()
	//pointMethod2()
	//pointMethod3()

	pointMethod5()
}

// 指针声明与初始化
func pointMethod1() {
	var p1 *int //指针零值为nil
	var p2 *string

	fmt.Println(p1)
	fmt.Println(p2)

	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2

	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(*p3)
	fmt.Println("----------------")
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}

// 使用指针访问值
func pointMethod2() {
	var p1 *int //声明一个指针p1
	i := 1      //局部变量i，通过赋值，类型推导确定i的类型
	p1 = &i     //指针p1指向i变量的地址

	fmt.Println(*p1 == i)

	*p1 = 8 //通过指针修改原始变量的值，i的值为1，通过指针修改为8
	fmt.Println(i)

	var p2 *string
	j := "test"
	p2 = &j
	fmt.Println(*p2 == j)
	fmt.Println(j)
}

// 修改指针指向的值
func pointMethod3() {
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}

func pointMethod4() {

	var a int = 100
	var p *int
	p = &a
	// p = p + 1    在 Go 中，指针不能直接参与计算，否则会在编译的时候就包错

	t := *p + 1 //通过获取指针的指向的值计算是可以的
	fmt.Println(t)
}

func pointMethod5() {
	var p *string
	var a string = "test"
	p = &a

	up1 := unsafe.Pointer(p)
	up2 := unsafe.Pointer(&a)

	fmt.Println(up1, up2)

	uintprNum := uintptr(up1)
	fmt.Println(uintprNum)
	uintprNum += 1
	fmt.Println(uintprNum)

	/*a := "Hello, world!"
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)*/
}
