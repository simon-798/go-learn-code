package lesson1

import "fmt"

/*
*
作业：
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func Job136Method(numArray []int) int {

	fmt.Println(numArray)

	countMap := make(map[int]int)

	for i := 0; i < len(numArray); i++ {
		v := numArray[i]
		value, ok := countMap[v]
		if !ok {
			countMap[v] = 1
		} else {
			countMap[v] = value + 1
		}
	}

	for i, v := range countMap {
		//fmt.Println(i, v)
		if v == 1 {
			return i
		}
	}

	return -1
}

// 第二种实现方法
func Job136Method2(numArray []int) int {

	fmt.Println(numArray)

	countMap := make(map[int]int)

	for _, v := range numArray {
		//Go的map底层是哈希表实现，每个键值对存储在桶里。当访问countMap[v]时，会先计算v的哈希值，找到对应的桶。
		//如果键存在，就返回对应的值；不存在的话，会返回零值。然后，++操作会将该值加1，如果是第一次出现，初始值0加1就变成了1
		//countMap[v]++	//等价于 countMap[v] = countMap[v] + 1
		countMap[v] = countMap[v] + 1
	}

	for k, v := range countMap {
		if v == 1 {
			return k
		}
	}

	return -1
}
