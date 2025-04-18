package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	curr := root

	for curr != nil {
		child := curr.Child

		if child != nil {
			next := curr.Next
			curr.Next = child
			child.Prev = curr

			for child.Next != nil {
				child = child.Next
			}

			curr.Child = nil
			child.Next = next
			if next != nil {
				next.Prev = child
			}
		}

		curr = curr.Next
	}

	return root
}
