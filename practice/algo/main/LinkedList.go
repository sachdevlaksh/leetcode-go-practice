package main

type LinkedList struct {
	head *ListNode
}

func (list LinkedList) Push(data int) {
	newNode := ListNode{Val: data}
	newNode.Next = list.head
	list.head = &newNode
}
