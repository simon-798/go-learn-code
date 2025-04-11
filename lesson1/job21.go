package main

import "fmt"

/*
*
 21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
    可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，
    将较小值的节点添加到新链表中，直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
*/
func job21Method() {
	// 测试用例
	//struct { list1 []int; list2 []int } 定义了一个未命名结构体类型，包含两个字段：
	//list1: 整型切片（[]int）
	//list2: 整型切片（[]int）
	testCases := []struct {
		list1 []int
		list2 []int
	}{
		{list1: []int{1, 3, 5}, list2: []int{2, 4, 6}}, // 合并结果: 1 -> 2 -> 3 -> 4 -> 5 -> 6
		{list1: []int{}, list2: []int{1, 2, 3}},        // 合并结果: 1 -> 2 -> 3
		{list1: []int{2, 4}, list2: []int{1, 3, 5}},    // 合并结果: 1 -> 2 -> 3 -> 4 -> 5
		{list1: []int{}, list2: []int{}},               // 合并结果: 空链表
	}

	for _, tc := range testCases {
		l1 := createList(tc.list1)
		l2 := createList(tc.list2)
		merged := mergeTwoLists(l1, l2)

		fmt.Print("合并结果: ")
		printList(merged)
	}
}

// 定义链表节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并2个升序链表，双指针参数l1,l2
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建虚拟头节点，简化链表头处理
	head := &ListNode{}
	currentPoint := head // 当前指针，用于构建新链表

	// 双指针遍历两个链表
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val { //如果l1链表的元素比l2链表的元素小
			currentPoint.Next = l1 // 将 l1 的当前节点接入新链表
			l1 = l1.Next           // l1 指针后移
		} else {
			currentPoint.Next = l2 // 将 l2 的当前节点接入新链表
			l2 = l2.Next           // l2 指针后移
		}
		currentPoint = currentPoint.Next // 移动新链表的当前指针
	}

	// l1链表和l2链表其中会有一个链表先循环完
	// 这里就用来处理剩余节点（直接拼接未遍历完的链表）
	if l1 != nil {
		currentPoint.Next = l1
	} else {
		currentPoint.Next = l2
	}

	// 返回合并后的链表头（虚拟头节点的下一个节点，虚拟头结点是类型的0值，跳过这个节点）
	return head.Next
}

// 辅助函数：将数组转换为链表
func createList(arr []int) *ListNode {
	head := &ListNode{}
	currentPoint := head //指针变量指向头节点
	for _, val := range arr {
		currentPoint.Next = &ListNode{Val: val} //指针为引用类型，所以把新的节点赋值指针变量的next，head节点的next也会跟着改变
		currentPoint = currentPoint.Next        //移动指针变量到next节点
	}
	return head.Next
}

// 辅助函数：打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

/*func main() {

	array := []struct {
		list1 []int
		list2 []int
	}{
		{list1: []int{1, 3, 5}, list2: []int{2, 4, 6}},
	}

	for _, v := range array {
		l1 := createList(v.list1)
		l2 := createList(v.list2)

		merged := mergeTwoLists(l1, l2)
		fmt.Println("合并结果：", merged)
	}

}*/
