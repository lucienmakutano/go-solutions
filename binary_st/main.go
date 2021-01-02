package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func main() {
	bt := &BST{}
	
	bt.insert(0).insert(-2).insert(2).insert(-1).insert(1).insert(-3).insert(3)
	fmt.Println(isValidBST(bt.root))
}

func (bt *BST) insert(val int) *BST {
	if bt.root == nil {
		bt.root = &Node{val: val, left: nil, right: nil}
	} else {
		bt.root.insert(val)
	}

	return bt

}

func (n *Node) insert(val int) {
	if n == nil {
		return
	} else if val <= n.val {
		if n.left == nil {
			n.left = &Node{val: val, left: nil, right: nil}
		} else {
			n.left.insert(val)
		}

	} else {
		if n.right == nil {
			n.right = &Node{val: val, left: nil, right: nil}
		} else {
			n.right.insert(val)
		}
	}

}

func isValidBST(n *Node) bool {

	if n.left == nil && n.right == nil {
		return true
	}

	if !(n.left.val < n.val) || !(n.right.val > n.val) {
		return false
	}

	return isValidBST(n.left) && isValidBST(n.right)
}
