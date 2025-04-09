package main

import "fmt"

func main() {
	/*	a, b := 1, 2
		sum := a + b
		sub := a - b
		mul := a * b
		div := a / b
		mod := a % b

		fmt.Println(sum, sub, mul, div, mod)*/

	/*a := 10 + 0.1
	b := byte(1) + 1
	fmt.Println(a, b)

	sum := a + float64(b)
	fmt.Println(sum)

	sub := byte(a) - b
	fmt.Println(sub)

	mul := byte(a) * b
	mul2 := a * float64(b)
	div := byte(a) / b
	div2 := a / float64(b)

	fmt.Println(mul, mul2, div, div2)

	fmt.Println(0 & 0)
	fmt.Println(0 | 0)
	fmt.Println(0 ^ 0)

	fmt.Println(0 & 1)
	fmt.Println(0 | 1)
	fmt.Println(0 ^ 1)

	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)

	fmt.Println(1 & 0)
	fmt.Println(1 | 0)
	fmt.Println(1 ^ 0)*/

	/*a, b := 1, 2
	var c int
	c = a + b
	fmt.Println("c = a + b, c =", c)

	plusAssignment(c, a)
	subAssignment(c, a)
	mulAssignment(c, a)
	divAssignment(c, a)
	modAssignment(c, a)
	leftMoveAssignment(c, a)
	rightMoveAssignment(c, a)
	andAssignment(c, a)
	orAssignment(c, a)
	norAssignment(c, a)*/

	/*a := 4
	var ptr *int
	fmt.Println(a)

	ptr = &a
	fmt.Printf("ptr 为 %d\n", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr)*/

	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	e = (a + b) * c / d // ( 31 * 16 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d // ( 31 * 16 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	e = (a + b) * (c / d) // 31 * (16/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	// 21 + (160/5)
	e = a + (b*c)/d
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)

	// 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4;4 | 12 = 12
	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println(f == 12)
}

func plusAssignment(c, a int) {
	c += a // c = c + a
	fmt.Println("c += a, c =", c)
}

func subAssignment(c, a int) {
	c -= a // c = c - a
	fmt.Println("c -= a, c =", c)
}

func mulAssignment(c, a int) {
	c *= a // c = c * a
	fmt.Println("c *= a, c =", c)
}

func divAssignment(c, a int) {
	c /= a // c = c / a
	fmt.Println("c /= a, c =", c)
}

func modAssignment(c, a int) {
	c %= a // c = c % a
	fmt.Println("c %= a, c =", c)
}

func leftMoveAssignment(c, a int) {
	c <<= a // c = c << a
	fmt.Println("c <<= a, c =", c)
}

func rightMoveAssignment(c, a int) {
	c >>= a // c = c >> a
	fmt.Println("c >>= a, c =", c)
}

func andAssignment(c, a int) {
	c &= a // c = c & a
	fmt.Println("c &= a, c =", c)
}

func orAssignment(c, a int) {
	c |= a // c = c | a
	fmt.Println("c |= a, c =", c)
}

func norAssignment(c, a int) {
	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c =", c)
}
