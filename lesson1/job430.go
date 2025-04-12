package main

import (
	"fmt"
)

/**
430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，
它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，
依此类推，生成多级数据结构，如下面的示例所示。给定位于列表第一级的头节点，请扁平化列表，
即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，
然后使用递归的方法来扁平化链表，先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
*/

func job430Method() {

	node := buildTestCaseNode()
	//list1 := linkedListToSlice(node)
	//fmt.Println("处理前:", list1)
	//将链表转换为数组，方便验证
	list2 := linkedListToSlice(flatten(node))
	fmt.Println("处理后:", list2)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node //子链表指针
}

/*
*
构建链表测试数据
*/
func buildTestCaseNode() *Node {

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	node9 := &Node{Val: 9}
	node10 := &Node{Val: 10}
	node11 := &Node{Val: 11}
	node12 := &Node{Val: 12}

	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2
	node3.Next = node4
	node4.Prev = node3
	node4.Next = node5
	node5.Prev = node4
	node5.Next = node6
	node6.Prev = node5

	node3.Child = node7

	node7.Next = node8
	node8.Next = node9
	node8.Prev = node7
	node9.Next = node10
	node9.Prev = node8
	node10.Prev = node9

	node8.Child = node11
	node11.Next = node12
	node12.Prev = node11

	return node1
}

/*
*

	单一子链表:
	1 <-> 2 <-> 3
	|
	v
	4 <-> 5
*/
func TestCase2() *Node {
	// 构建输入链表
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}

	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2

	node1.Child = node4
	node4.Next = node5
	node5.Prev = node4

	return node1

}

// 将链表转换为数组，方便验证
func linkedListToSlice(head *Node) []int {
	var slice []int
	curr := head
	for curr != nil {
		slice = append(slice, curr.Val)
		curr = curr.Next
	}
	return slice
}

func flatten(node *Node) *Node {

	if node == nil {
		return nil
	}

	flattenDFS(node)
	return node
}

/*
*

	深度优先遍历DFS：优先处理子链表，再处理后续节点。
	递归处理：递归展平子链表并返回尾节点，将其插入到当前节点与后续节点之间。
*/
func flattenDFS(node *Node) *Node {
	curr := node   //当前处理的节点
	var last *Node //记录当前链表的最后一个尾节点

	for curr != nil {
		next := curr.Next //保存原next节点

		if curr.Child != nil { //如果当前节点的子链表不为空
			//递归展平子链表，获取子链表的尾节点
			childTail := flattenDFS(curr.Child)

			//将子链表插入current和next之间
			curr.Next = curr.Child
			curr.Child.Prev = curr

			childTail.Next = next
			if next != nil {
				next.Prev = childTail
			}

			//清空child指针，避免循环引用
			curr.Child = nil

			//更新last为子链表的尾节点
			last = childTail
		} else {
			//没有子链表，当前节点就是最后一个尾节点
			last = curr
		}

		curr = next
	}

	return last
}
