package main

import "fmt"

/*
  题目链接：https://leetcode-cn.com/problems/add-two-numbers/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func addTwoNumbers0(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := NewListNode(0)
	cur := pre
	// 代表进位
	carry := 0
	for l1 != nil || l2 != nil{
		x,y := 0, 0
		if l1 == nil{
			x = 0
		}else {
			x = l1.Val
		}
		if l2 == nil{
			y = 0
		}else {
			y = l2.Val
		}
		sum := x + y + carry
		carry = sum/10
		val := sum%10
		cur.Val = val
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil{
			next := NewListNode(0)
			cur.Next = next
			cur = next
		}
	}
	if carry != 0{
		next := NewListNode(carry)
		cur.Next= next

	}

	return pre
}

func main()  {
	res1 := NewListNode(1)
	res2 := NewListNode(2)
	fmt.Println(addTwoNumbers0(res1, res2))
}