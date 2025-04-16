package lesson1

import "fmt"

/*
*
729. 我的日程安排表 I：
实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成重复预订，
则可以存储这个新的日程安排。当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），
就会产生 重复预订 。
日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end) 即包含 start 但不包含 end，
实数 x 的范围为 start <= x < end 。实现 MyCalendar 类：MyCalendar() 初始化日历对象。
boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true ，
否则，返回 false 并且不要将该日程安排添加到日历中。
可以定义一个结构体来表示日程安排，包含 start 和 end 字段，然后使用一个切片来存储所有的日程安排，
在 book 方法中，遍历切片中的日程安排，判断是否与要添加的日程安排有重叠。
*/
func Job729Method() {

	//实例化MyCalendar 结构体和events二维数组
	calendar := MyCalendar{events: [][]int{}}
	fmt.Println(calendar.Book(10, 20)) // true
	fmt.Println(calendar.Book(15, 25)) // false
	fmt.Println(calendar.Book(20, 30)) // true

}

/*
*
// MyCalendar 结构体，存储所有已预订的区间
*/
type MyCalendar struct {
	events [][]int // 每个元素是 [start, end)

}

/*
*
Book方法尝试添加新的时间区间，返回是否成功
*/
func (calendar *MyCalendar) Book(start, end int) bool {

	//第一个区间，不存在重叠，直接设置成功
	if calendar.events == nil || len(calendar.events) == 0 {
		newEvents := []int{start, end}
		calendar.events = append(calendar.events, newEvents)
		return true
	}

	flag := false
	// [1,3],[2,5]
	for _, event := range calendar.events {
		// 检查是否重叠：新区间start < 旧区间end 且 新区间end > 旧区间start
		if start < event[1] && end > event[0] {
			return flag
		}

		calendar.events = append(calendar.events, []int{start, end})
		flag = true
		break
	}

	return flag

}
