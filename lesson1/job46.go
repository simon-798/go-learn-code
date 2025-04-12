package main

import "fmt"

/**
46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
	可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，
	使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
*/

// 没有完全理解回溯算法
func job46Method() [][]int {
	nums := []int{1, 2}

	var result [][]int                    //存储所有排列的结果  二维切片
	var track []int                       //用于记录当前已经选择的数字  一维切片
	used := make([]bool, len(nums))       //用于标记是否已经被使用过 长度与nums相同
	backtrack(nums, track, used, &result) //开始回溯过程
	fmt.Println("result:", result)
	return result
}

func backtrack(nums []int, track []int, used []bool, result *[][]int) {

	// 结束条件：当路径长度等于数组长度时，说明找到了一种排列
	if len(track) == len(nums) {
		// 注意这里需要复制一份 track，因为后续 track 会被修改
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		return
	}

	//
	for i := 0; i < len(nums); i++ {
		//是否选择过
		if used[i] {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		used[i] = true
		// 递归进入下一层
		backtrack(nums, track, used, result)

		// 撤销选择
		last := len(track) - 1
		//从开始下标位置0，到结束下标位置last的元素，创建一个新的切片，相对于移除最后一个元素
		track = track[:last]

		used[i] = false
	}
}
