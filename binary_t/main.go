package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BTree struct {
	Root *Node
}

func main() {
	s := &BTree{}
	s.Insert(1).Insert(3).Insert(8)
	t := &Node{}
	t = newNode(1)
	t.Left = nil
	t.Right = newNode(8)

	fmt.Println(IsSubTree(s.Root, t))
}

func IsSubTree(s, t *Node) bool {
	

	return true
}

func (bt *BTree) Insert(val int) *BTree {
	if bt.Root == nil {
		bt.Root = newNode(val)
	} else {
		bt.Root.Insert(val)
	}

	return bt
}

func (r *Node) Insert(val int) {
	if r == nil {
		return
	}

	q := list.New()
	q.PushBack(r)

	for q.Len() > 0 {
		el := q.Front()
		value := el.Value
		q.Remove(el)
		n, ok := value.(*Node)
		if !ok {
			break
		}

		if n.Left == nil {
			n.Left = newNode(val)
			break
		} else {
			q.PushBack(n.Left)
		}

		if n.Right == nil {
			n.Right = newNode(val)
			break
		} else {
			q.PushBack(n.Right)
		}
	}
}

func newNode(val int) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}

func Print(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	Print(root.Left)
	Print(root.Right)
}
