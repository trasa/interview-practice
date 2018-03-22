package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func newBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{
		value: value,
	}
}

func (n *BinaryTreeNode) insertLeft(leftVal int) *BinaryTreeNode {
	n.left = newBinaryTreeNode(leftVal)
	return n.left
}

func (n *BinaryTreeNode) insertRight(rightVal int) *BinaryTreeNode {
	n.right = newBinaryTreeNode(rightVal)
	return n.right
}

func TestSecondLargestBinarySearch(t *testing.T) {
	root := newBinaryTreeNode(10)
	root.insertLeft(5).insertRight(6)
	r := root.insertRight(20)
	r.insertLeft(17)
	r.insertRight(30).insertRight(40).insertLeft(35)

	second := findSecondLargest(root)
	assert.Equal(t, 35, second.value)
	fmt.Printf("second is %d\n", second.value)
}

func findSecondLargest(root *BinaryTreeNode) *BinaryTreeNode {

	for cur := root; cur != nil; cur = cur.right {
		// if cur is largest and has a left subtree,
		// then 2nd largest is the largest in the subtree
		if cur.left != nil && cur.right == nil {
			return findLargest(cur.left)
		}

		// current is parent of largest, largest has no children
		// ... current = 2nd largest
		if cur.right != nil && cur.right.left == nil && cur.right.right == nil {
			return cur
		}
		// otherwise keep looking to the right
	}
	return nil
}

func findLargest(n *BinaryTreeNode) *BinaryTreeNode {
	for ; n.right != nil; n = n.right {
	}
	return n
}
