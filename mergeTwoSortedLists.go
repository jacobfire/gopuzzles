package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//buffer := ListNode{Val: 0, Next: nil}
	//resultNode := ListNode{Val: 0, Next: nil}
	//var flag int = 1
	//var flagInternal int = 1
	//var sum int
	//var currentNode, currentSecondNode *ListNode
	//
	//currentNode = l1
	//currentSecondNode = l2
	//for flag > 0 {
	//	sum += currentNode.Val
	//	if currentNode.Next == nil {
	//		fmt.Print("flag = 0")
	//		flag = 0
	//		//fmt.Print(l1)
	//		fmt.Print("currentNode = ")
	//		fmt.Print(currentNode)
	//	} else {
	//		//for flagInternal > 0 {
	//		//	if currentSecondNode.Next == nil {
	//		//		flagInternal = 0
	//		//	}
	//		//	if currentSecondNode.Val >= currentNode.Val &&  currentSecondNode.Val < currentNode.Next.Val {
	//		//		fmt.Print("condition A")
	//		//		//if currentSecondNode.Val == 1 {
	//		//	//	fmt.Print(currentNode.Val)
	//		//	//	fmt.Print(currentSecondNode.Val)
	//		//	//	fmt.Print(currentSecondNode.Next.Val)
	//		//		buffer = buffer
	//		//		//buffer.Next = currentNode.Next
	//		//		link := currentNode.Next
	//		//		currentSecondNode.Next = link
	//		//		currentNode.Next = currentSecondNode
	//		//
	//		//	} else {
	//		//		fmt.Print("condition B")
	//		//		fmt.Println("currentSecondNode = ")
	//		//		fmt.Println(currentSecondNode.Val)
	//		//		currentSecondNode = currentSecondNode.Next
	//		//	}
	//		//
	//		//}
	//		buffer = buffer
	//		flagInternal = flagInternal
	//		startElementInSecondNode := l2
	//		if currentSecondNode.Val >= currentNode.Val &&  startElementInSecondNode.Val < currentNode.Next.Val {
	//			fmt.Print("Start = ")
	//			fmt.Print(startElementInSecondNode.Val)
	//			nextSecondElement := startElementInSecondNode.Next
	//			startElementInSecondNode.Next = currentNode.Next
	//			currentNode.Next = startElementInSecondNode
	//
	//			l2 = nextSecondElement
	//
	//		} else {
	//			//if currentNode.Next
	//			fmt.Print("; l2.Val = ")
	//			fmt.Print(l2.Val)
	//			if currentNode.Next == nil && l2 != nil {
	//				link := currentNode.Next
	//				if l2.Next == nil {
	//					l2 = nil
	//				} else {
	//					l2 = l2.Next
	//					currentSecond := l2
	//					currentSecond.Next = link
	//				}
	//			}
	//
	//		}
	//
	//		currentNode = currentNode.Next
	//	}
	//}
	//
	//flag = 1
	////sum = 0
	//////currentNode = l1
	////for flag > 0 {
	////	if currentNode.Next == nil {
	////		flag = 0
	////	} else {
	////		//sum += currentNode.Val
	////		sum ++
	////		currentNode = currentNode.Next
	////	}
	////}
	//resultNode.Val = sum
	//return &resultNode
	node := &ListNode{}
	root := node

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = &ListNode{Val: l1.Val}
			node = node.Next
			l1 = l1.Next
		} else {
			node.Next = &ListNode{Val: l2.Val}
			node = node.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return root.Next
}

func main() {
	//var l1, l2 ListNode
	var l11, l12, l13, l21, l22, l23 ListNode
	//+++++++++++++++++++++++
	l13.Val = 3
	l13.Next = nil

	l12.Val = 2
	l12.Next = &l13

	l11.Val = 1
	l11.Next = &l12
	//+++++++++++++++++++++++
	l23.Val = 4
	l23.Next = nil

	l22.Val = 3
	l22.Next = &l23

	l21.Val = 1
	l21.Next = &l22
	//+++++++++++++++++++++++

	result := mergeTwoLists(&l11, &l21)

	fmt.Print(result)
}
