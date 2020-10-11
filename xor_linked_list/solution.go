package xor_linked_list

// XORLinkedList is an implementation of a doubly Linked list using xor
type XORLinkedList struct {
	El   int
	Prev *XORLinkedList
	Next *XORLinkedList
}

// Add adds an element to the list
func (xor *XORLinkedList) Add(el int) {
	var newNode *XORLinkedList = &XORLinkedList{}

	if xor.Prev.El == 0 {
		xor.El = el
		xor.Next = newNode
		return
	}

	newNode.El = el
	newNode.Prev = xor
	xor.Next = newNode
	newNode.Next = &XORLinkedList{}
}

// Get retrieves an element fro the list based on a index
func (xor *XORLinkedList) Get(idx int) {}
