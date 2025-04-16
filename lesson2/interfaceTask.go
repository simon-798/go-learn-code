package lesson2

import (
	"fmt"
	"math"
)

/*
*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
func InterfaceTask1() {

	rectangle := Rectangle{width: 10, height: 5}
	circle := Circle{Radius: 7}

	shapes := []Shape{&rectangle, &circle}
	for index, shape := range shapes {
		fmt.Printf("形状 %d:\n", index+1)
		fmt.Printf("面积: %.2f\n", shape.Area())
		fmt.Printf("周长: %.2f\n", shape.Perimeter())
		fmt.Println("------------------")
	}

}

// 定义形状接口
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// 矩形结构体
type Rectangle struct {
	height float64 //高
	width  float64 //宽
}

// 实现矩形的面积计算
func (rectangle *Rectangle) Area() float64 {

	return rectangle.height * rectangle.width
}

// 实现矩形的周长计算
func (rectangle *Rectangle) Perimeter() float64 {
	return 2 * (rectangle.height + rectangle.width)
}

// 圆形结构体
type Circle struct {
	Radius float64 //半径
}

// 实现圆形的面积计算
func (circle *Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

/*
*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
func InterfaceTask2() {
	person := Person{
		Name: "simon",
		Age:  18,
	}

	emp := &Employee{Person: person, EmployeeID: "1001"}
	emp.PrintInfo()

}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (employee *Employee) PrintInfo() {
	fmt.Printf("EmployeeID:%s,Name:%s, Age:%d\n", employee.EmployeeID, employee.Name, employee.Age)
}
