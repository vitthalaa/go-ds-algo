package main

// BinaryTree This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	_, isBalanced := isBalancedHeight(tree, 0)

	return isBalanced
}

func isBalancedHeight(node *BinaryTree, height int) (int, bool) {
	if node == nil {
		return height, true
	}

	left, isLeftBalanced := isBalancedHeight(node.Left, height+1)
	right, isRightBalanced := isBalancedHeight(node.Right, height+1)

	if left >= right {
		return left, isLeftBalanced && isRightBalanced && left <= right+1
	}

	return right, isLeftBalanced && isRightBalanced && right <= left+1
}
