package basics

import "fmt"

type Data struct{ ID int }

/*
*
切片
*/
func mainSliceTest() {
	/*a := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	a[0] = 9
	a[1] = 8
	a[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	fmt.Println(a)*/

	//sliceTest2()
	//sliceTest3()
	sliceTest8()

	//sliceTest6()

	/*slice := []*Data{{1}, {2}, {3}}
	fmt.Println("slice1:", slice)

	copy(slice[0:], slice[1:])
	for i := 0; i < len(slice); i++ {
		fmt.Println("for :", i, *slice[i], slice[i].ID)
	}

	slice[len(slice)-1] = nil // 帮助 GC 回收
	fmt.Println("slice3:", slice[0].ID, slice[1].ID, len(slice))
	for i := 0; i < len(slice)-1; i++ {
		fmt.Println(" for2 :", i, slice[i], slice[i].ID)
	}

	slice = slice[:len(slice)-1]
	//fmt.Println("slice4:", slice[0].ID, slice[1].ID)

	for i := 0; i < len(slice); i++ {
		fmt.Println("gc for :", i, slice[i], slice[i].ID)
	}*/
}

func sliceTest2() {
	s1 := []int{5, 4, 3, 2, 1}
	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(e1, e2, e3)
	fmt.Println(s1)

	// 向指定位置赋值
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Println(s1)

	// range迭代访问切片
	for i, v := range s1 {
		fmt.Println("before modify, s1[%d] = %d", i, v)
	}

}

func sliceTest3() {
	var nilSlice []int
	fmt.Println("nilSlice length:", len(nilSlice))
	fmt.Println("nilSlice capacity:", len(nilSlice))

	s2 := []int{9, 8, 7, 6, 5}
	fmt.Println("s2 length: ", len(s2))
	fmt.Println("s2 capacity: ", cap(s2))
}

func sliceTest4() {
	s3 := []int{}
	fmt.Println("s3 = ", s3)

	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	s4 := append(s3, 10)
	fmt.Println("s3 = ", s3)
	fmt.Println("s4 = ", s4)

	s6 := make([]int, 2, 4)
	fmt.Println(s6)

	s6 = append(s6)
	s6 = append(s6, 1)
	s6 = append(s6, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s7 := append(s6, 11)
	fmt.Println(s6)
	fmt.Println(s7)
}

func sliceTest5() {
	/*s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{108}, s4[2:]...)...)
	fmt.Println("s4 = ", s4)*/

	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("1.s5 = ", s5)

	s5 = s5[:3] // 截取下标3之前的所有元素
	fmt.Println("2.s5 = ", s5)
	s5 = s5[2:]
	fmt.Println("3.s5 = ", s5)
	s5 = s5[0:3]
	fmt.Println("4.s5 = ", s5)

}

func sliceTest6() {

	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}
func sliceTest7() {

	s := make([]int, 3, 6)
	fmt.Println("s length:", len(s))
	fmt.Println("s capacity:", cap(s))
	fmt.Println("initial, s = ", s)
	s[1] = 2
	fmt.Println("set position 1, s = ", s)

	modifySlice(s)
	fmt.Println("after modifySlice, s = ", s)

}

func modifySlice(param []int) {
	param[0] = 1024
}

func sliceTest8() {
	s := make([]int, 3, 6)
	fmt.Println("initial, s =", s)
	s[1] = 2
	fmt.Println("after set position 1, s =", s)

	s2 := append(s, 4)
	fmt.Println("after append, s length:", len(s))
	fmt.Println("after append, s capacity:", cap(s))
	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s =", s)
	fmt.Println("after append, s2 =", s2)

	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	fmt.Println("after set position 0, s2 =", s2)

	appendInFunc(s)
	fmt.Println("after append in func, s =", s)
	fmt.Println("after append in func, s2 =", s2)
}

func appendInFunc(param []int) {
	param = append(param, 1022)
	param = append(param, 1023)
	fmt.Println("in func, param =", param)
	param[2] = 512
	fmt.Println("2.in func, param =", param)
	fmt.Println("set position 2 in func, param =", param)
}
