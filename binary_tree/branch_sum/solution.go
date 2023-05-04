package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sum := []int{}
	calculateBranchSum(root, 0, &sum)
	return sum
}

func calculateBranchSum(node *BinaryTree, runningSum int, sum *[]int) {
	if node == nil {
		return
	}

	runningSum = runningSum + node.Value
	if node.Right == nil && node.Left == nil {
		*sum = append(*sum, runningSum)
		return
	}

	calculateBranchSum(node.Left, runningSum, sum)
	calculateBranchSum(node.Right, runningSum, sum)
}

func main() {
	/*
			10
		   / \
		  5   12
		 /     \
		8 		16
	*/
	tree := &BinaryTree{Value: 10}
	tree.insert(5)
	tree.insert(12)
	tree.insert(16)
	tree.insert(8)

	sums := BranchSums(tree)

	fmt.Println(sums)
}

// helper function to construct binary tree for verifying solution
func (t *BinaryTree) insert(val int) {
	if t == nil {
		return
	}

	if val <= t.Value {
		if t.Left == nil {
			t.Left = &BinaryTree{Value: val}
			return
		}

		t.Left.insert(val)
		return
	}

	if t.Right == nil {
		t.Right = &BinaryTree{Value: val}
		return
	}

	t.Right.insert(val)
	return
}
