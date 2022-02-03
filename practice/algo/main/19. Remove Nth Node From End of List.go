package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	current := head
	var prev *ListNode
	for head != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next

	}
	temp := head
	i := 0
	for i < n-1 && temp != nil {
		temp = temp.Next
		if temp == nil || temp.Next == nil {
			return nil
		}
		next := temp.Next.Next
		temp.Next = next
		i++
	}

	return head
}
