package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMaxPathSum(node *TreeNode, maxPathSum *int) int {
	if node == nil {
		return 0
	}

	leftMax := findMaxPathSum(node.Left, maxPathSum)
	rightMax := findMaxPathSum(node.Right, maxPathSum)

	pathSum := node.Val
	if leftMax > 0 {
		pathSum += leftMax
	}
	if rightMax > 0 {
		pathSum += rightMax
	}
	if pathSum > *maxPathSum {
		*maxPathSum = pathSum
	}

	maxBranchSum := 0
	if leftMax > maxBranchSum {
		maxBranchSum = leftMax
	}
	if rightMax > maxBranchSum {
		maxBranchSum = rightMax
	}

	return maxBranchSum + node.Val
}

func maxPathSum(root *TreeNode) int {
	maximumPathSum := root.Val
	findMaxPathSum(root, &maximumPathSum)
	return maximumPathSum
}
