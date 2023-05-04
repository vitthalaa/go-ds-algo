package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return calculateDepth(root, 0)
}

func calculateDepth(node *BinaryTree, depth int) int {
	if node == nil {
		return depth
	}

	right := 0
	if node.Right != nil {
		right = calculateDepth(node.Right, depth+1)
	}

	left := 0
	if node.Left != nil {
		left = calculateDepth(node.Left, depth+1)
	}

	return depth + left + right
}

func main() {
	/*
				10
			   / \
			  5   12
			 /\   /\
			3  8 11	16
		   /\
		  1  4
	*/
	tree := &BinaryTree{Value: 10}
	tree.insert(5)
	tree.insert(12)
	tree.insert(11)
	tree.insert(16)
	tree.insert(8)
	tree.insert(3)
	tree.insert(1)
	tree.insert(4)

	depth := NodeDepths(tree)

	fmt.Println(depth)
}

// helper function to construct binary tree for verifying solution
func (t *BinaryTree) insert(val int) {
	if t == nil {
		return
	}

	if val <= t.Value {
		if t.Left == nil {
			t.Left = &BinaryTree{Value: val, Left: nil, Right: nil}
			return
		}

		t.Left.insert(val)
		return
	}

	if t.Right == nil {
		t.Right = &BinaryTree{Value: val, Left: nil, Right: nil}
		return
	}

	t.Right.insert(val)
	return
}
