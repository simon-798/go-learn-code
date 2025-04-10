package main

/*
*
作业：
198. 打家劫舍：你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，
系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。
这道题可以使用动态规划的思想，通过 for 循环遍历数组，利用 if 条件判断来决定是否选择当前房屋进行抢劫，
状态转移方程为 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])。
*/
func Job198Method(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	//dp := make([]int, len(nums))

	perv := nums[0]                  //2
	current := max(nums[0], nums[1]) //7
	//从第三个元素开始遍历
	for i := 2; i < len(nums); i++ {

		//dp[i] = max(dp[i-1],dp[i-2] + nums[i])

		temp := max(current, perv+nums[i])
		perv = current
		current = temp

	}
	return current
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
