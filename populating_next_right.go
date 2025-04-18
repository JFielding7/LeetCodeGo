package main

//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}
//
//func connect(root *Node) *Node {
//	if root == nil || root.Left == nil {
//		return root
//	}
//
//	parent := root
//	head := root.Left
//	for head != nil {
//		prev := parent.Right
//		parent.Left.Next = prev
//		parent = parent.Next
//
//		for parent != nil {
//			prev.Next = parent.Left
//			parent.Left.Next = parent.Right
//			prev = parent.Right
//			parent = parent.Next
//		}
//
//		parent = head
//		head = head.Left
//	}
//
//	return root
//}
