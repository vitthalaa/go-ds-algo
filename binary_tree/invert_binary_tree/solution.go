package main

import "fmt"

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// InvertBinaryTree O(n) ST
func (tree *BinaryTree) InvertBinaryTree() {
	invertNode(tree)
}

func invertNode(node *BinaryTree) {
	if node == nil {
		return
	}

	t := node.Left
	node.Left = node.Right
	node.Right = t

	invertNode(node.Left)
	invertNode(node.Right)
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

	tree.InvertBinaryTree()

	fmt.Println(tree)
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
