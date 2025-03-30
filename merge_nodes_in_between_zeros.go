package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	curr := head.Next

	for curr != nil {
		sum := 0
		for curr.Val != 0 {
			sum += curr.Val
			curr = curr.Next
		}

		newTail := &ListNode{Val: sum}
		tail.Next = newTail
		tail = newTail
		curr = curr.Next
	}

	return dummy.Next
}
