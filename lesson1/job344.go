package main

import "fmt"

/*
*
 344. 反转字符串：编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
    不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。可以使用 for 循环和两个指针，
    一个指向字符串的开头，一个指向字符串的结尾，然后交换两个指针所指向的字符，直到两个指针相遇。
*/
func job344Method(strArray []string) {

	if len(strArray) == 0 {
		fmt.Println("数组为空")
		return
	}
	fmt.Println("反转前:", strArray)

	//将 left 指向字符数组首元素
	left := 0
	//right 指向字符数组尾元素
	right := len(strArray) - 1

	for ; left < right; left++ {
		temp := strArray[left]
		strArray[left] = strArray[right]
		strArray[right] = temp
		right--
	}
	fmt.Println("反转后:", strArray)
}

/*
*
写法2
*/
func job344Method2() {

	strArray := []string{"a", "b", "c", "d", "e", "f"}

	if len(strArray) == 0 {
		fmt.Println("数组为空")
		return
	}
	fmt.Println("反转前:", strArray)

	//将 left 指向字符数组首元素
	//right 指向字符数组尾元素
	for left, right := 0, len(strArray)-1; left < right; left++ {
		strArray[left], strArray[right] = strArray[right], strArray[left]
		right--
	}
	fmt.Println("反转后:", strArray)
}
