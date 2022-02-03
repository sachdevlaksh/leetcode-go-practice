package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Insert(val int) {
	n := ListNode{}
	n.Val = val
	slow:=l
	count := 0
	for slow != nil && slow.Next == nil {
		count++
		slow = slow.Next
		}
			
	if count == 0 {
		l.Val = val
		
	}else{	
	for i := 0; i < count; i++ {
		if l.Next == nil {
			l.Next = &n			
			return
		}
		l = l.Next
	}
}
}
func middleNode(head *ListNode) *ListNode {
	nodeCount := 0
	headNodeCopy := head
	for headNodeCopy.Next != nil {
		nodeCount++
		headNodeCopy = headNodeCopy.Next
	}
	mid := 0
	if nodeCount%2 == 0 {
		mid = nodeCount / 2
	} else {
		mid = nodeCount/2 + 1
	}
	fmt.Printf("mid %d ", mid)
	count := 0
	for head.Next != nil {
		fmt.Println(head)
		head = head.Next
		count++
		if count == mid {
			break
		}
	}
	outNode := head
	fmt.Print(outNode)
	return outNode
}
