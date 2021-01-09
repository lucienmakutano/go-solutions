package binaryst

import "fmt"

// Node represents a child node in a binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST represents a binary tree data structure
type BST struct {
	Root *Node
}

// NewBST creates a empty binary search tree
func NewBST() *BST {
	return &BST{}
}

// Insert implements insertion in a binary tree
func (bt *BST) Insert(val int) *BST {
	if bt.Root == nil {
		bt.Root = &Node{Value: val, Left: nil, Right: nil}
	} else {
		bt.Root.Add(val)
	}
	return bt
}

// Add creates a new node and appends it to the tree
func (root *Node) Add(val int) {
	//if n == nil {
	//	return
	//} else if val <= n.Value {
	//	n.Left == nil {
	//		n.Left = &Node {Value: val, Left: nil, Right: nil}
	//	} else {
	//		n.Left.Add(val)
	//	}
	//} else {
	//	n.Right == nil {
	//		n.Right = &Node {Value: val, Left: nil, Right: nil}
	//	} else {
	// 		n.Right.Add(val)
	//	}
	// }

	if root == nil {
		return
	}

	for {
		if val <= root.Value {
			root = root.Left

			if root == nil {
				root = &Node{Value: val, Left: nil, Right: nil}
				break
			}
		} else {
			root = root.Right

			if root == nil {
				root = &Node{Value: val, Left: nil, Right: nil}
				break
			}
		}
	}
}

// Find implements search in a binary search tree
func (root *Node) Find(val int) *Node {

	if root.Value == val {
		return root
	}

	if root.Left == nil && root.Right == nil {
		return nil
	}

	if val < root.Value {
		return root.Left.Find(val)
	}
	return root.Right.Find(val)
}

// PrintTree prints the tree
func PrintTree(root *Node) {
	if root == nil {
		return
	}

	PrintTree(root.Left)

	PrintTree(root.Right)

	fmt.Println(root.Value)
}

// IsValidBST checks if the passed tree is a valid binary search tree
func IsValidBST(root *Node) bool {

	if root.Left == nil && root.Right == nil {
		return true
	}

	if !(root.Left.Value < root.Value) || !(root.Right.Value > root.Value) {
		return false
	}

	return IsValidBST(root.Left) && IsValidBST(root.Right)
}
