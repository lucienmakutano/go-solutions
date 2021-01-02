package main

import (
	"fmt"
)

type Node struct{
	val int
	left *Node
	right *Node
}

type BT struct {
	root *Node
}

func main() {
	
	bTree := &Node {}
	
	bTree.Add(1)
	bTree.Add(3)
	bTree.Add(2)
	bTree.Add(5)
	bTree.Add(-1)
	bTree.Add(-2)
	bTree.Add(-4)
	
	// printNodeValue(bTree)
	fmt.Println(Find(bTree, 5).val)
}

func (bt *BT) insert(val int) *BT{
	if bt.root == nil {
		bt.root = &Node{val: val, left: nil, right: nil,}
	}
	else {
		bt.root.Add(val)
	}
	return bt
}


func (n *Node) Add(val int) {
	//if n == nil {
	//	return
	//} else if val <= n.val {
	//	n.left == nil {
	//		n.left = &Node {val: val, left: nil, right: nil}
	//	} else {
	//		n.left.Add(val)
	//	}
	//} else {
	//	n.right == nil {
	//		n.right = &Node {val: val, left: nil, right: nil}
	//	} else {
	// 		n.right.Add(val)
	//	}
	// }

	if n == nil {
		return
	}
	
	for {
		if val <= root.val {
			root = root.left
			
			if root == nil {
				root = &Node {val: val, left: nil, right: nil,}
				break
			}
		} else {
			root = root.right
			
			if root == nil {
				root = &Node {val: val, left: nil, right: nil,}
				break
			}
		}
	}
}


func (r *Node) Find(val int) *Node {

	if r.val == val {
		return r
	}

	if r.left == nil && r.right == nil {
		return nil
	}
	
	if val < r.val {
		return r.left.Find(val)
	} else {
		return r.right.Find(val)
	}
}

func printNodeValue(node *Node){
	if node == nil {
		return
	}
	
	
	printNodeValue(node.left)
	
	printNodeValue(node.right)
	
	fmt.Println(node.val)
}
