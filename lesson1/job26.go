package main

import "fmt"

/*
*
 26. 删除有序数组中的重复项：
    给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
    不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
    可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
    当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func job26Method() {

	nums := [][]int{
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{2, 2, 2},
		{1, 2, 3, 4, 5},
		{1, 1, 2},
	}

	for _, v := range nums {
		removeDuplicates(v)
	}
}

func removeDuplicates(nums []int) {
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	i += 1

	nums = append(nums[0:i])
	fmt.Println("去重后的数据:", nums, "，去重后数组的新长度:", i)
}
