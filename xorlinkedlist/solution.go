package xorlinkedlist

import "fmt"

// XORLinkedList is an implementation of a doubly Linked list using xor
type XORLinkedList struct {
	El   int
	Prev *XORLinkedList
	Next *XORLinkedList
}

// New return XORLinkedList type
func New() *XORLinkedList {
	return &XORLinkedList{}
}

// Insert adds an element to the list
func (xor *XORLinkedList) Insert(el int) {
	if xor == nil {
		xor = &XORLinkedList{El: el, Prev: nil, Next: nil}
		return
	}

	for xor != nil {

		if xor.Next == nil {
			xor.Next = &XORLinkedList{El: el, Prev: xor, Next: nil}
			break
		}

		xor = xor.Next
	}
}

// Get returns an element from the list based on a index
func (xor *XORLinkedList) Get(idx int) *XORLinkedList {
	counter := 0

	for xor != nil {

		if counter == idx {
			return xor
		}

		counter++
		xor = xor.Next
	}

	return nil
}

// PrintList prints all node values in the list
func PrintList(xor *XORLinkedList) {
	for xor != nil {
		fmt.Println(xor.El)
		xor = xor.Next
	}
}
