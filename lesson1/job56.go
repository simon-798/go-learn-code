package lesson1

import (
	"fmt"
	"sort"
)

/**
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/

func Job56Method() {

	testCase := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{2, 3}, {1, 4}},
	}

	for _, row := range testCase {
		fmt.Println("合并前:", row)
		merged := merge(row)
		fmt.Println("合并后:", merged)
	}

}

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return [][]int{}
	}

	//intervals二维数组按区间的start元素升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Println("sort:",intervals)

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i] //当前循环的区间数组

		last := merged[len(merged)-1] //取合并的结果集中的最后一个区间集合（结果集已经按升序排序）

		if current[0] <= last[1] { //合并
			merged[len(merged)-1][1] = max(last[1], current[1])

		} else {
			merged = append(merged, current)
		}
	}

	return merged

}
