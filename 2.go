package main

import "fmt"

// 2. 两数相加
// 难度 中等
// https://leetcode-cn.com/problems/add-two-numbers/

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1i []int = []int{2, 4, 3}
	var l2i []int = []int{5, 6, 4}
	//var l1i []int = []int{0}
	//var l2i []int = []int{0}
	//var l1i []int = []int{9, 9, 9, 9, 9, 9, 9}
	//var l2i []int = []int{9, 9, 9, 9}
	var l2t *ListNode = &ListNode{}
	var l2 *ListNode = l2t
	for _, v := range l2i {
		a := l2t
		a.Val = v
		a.Next = &ListNode{0, nil}
		l2t = a.Next
	}

	var l1t *ListNode = &ListNode{}
	var l1 *ListNode = l1t
	for _, v := range l1i {
		a := l1t
		a.Val = v
		a.Next = &ListNode{0, nil}
		l1t = a.Next
	}

	r := addTwoNumbers(l1, l2)
	for {
		if r.Next != nil {
			fmt.Print(r.Val)
			fmt.Print(" ")
			r = r.Next
			continue
		}
		break
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode = &ListNode{}
	tempResult := result
	var nextVal = 0
	for {

		var l1Val, l2Val int = 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		var l12Val int = l1Val + l2Val + nextVal
		a := l12Val % 10
		nextVal = l12Val / 10
		result.Val = a

		if l1 == nil && l2 == nil && l12Val == 0 {
			result.Next = nil
			break
		} else {
			nextListNode := ListNode{}
			result.Next = &nextListNode
			result = &nextListNode
		}
	}
	return tempResult
}
