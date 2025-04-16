package lesson1

import "fmt"

/**
69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。可以使用二分查找法来解决，
定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，
直到找到满足条件的平方根或者确定不存在精确的平方根。
*/

func Job69Method() {

	nums := []int{8, 9}
	for _, x := range nums {
		answer := sqrt(x)
		fmt.Println("计算", x, "的平方根为:", answer)
	}

}

func sqrt(x int) int {
	if x <= 0 { //特殊情况处理：x=0的平方根为0
		return 0
	}

	left := 1  // 左边界从1开始（因为x>=1时，平方根至少为1）
	right := x // 右边界初始为x
	ans := 0   // 保存最终结果

	for left <= right {
		mid := left + (right-left)/2 // 计算中间值，避免整数溢出
		midVal := mid

		// 用除法代替乘法，防止 mid*mid 溢出
		if midVal <= x/midVal && midVal*midVal <= x {
			ans = mid      // 当前 mid 可能是候选答案
			left = mid + 1 // 尝试更大的值
		} else {
			right = mid - 1 // 调整右边界，缩小范围
		}
	}

	return ans
}
