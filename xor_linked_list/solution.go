package xor_linked_list

// XORLinkedList is an implementation of a doubly Linked list using xor
type XORLinkedList struct {
	El   int
	Prev *XORLinkedList
	Next *XORLinkedList
}

// Add adds an element to the list
func (xor *XORLinkedList) Add(el int) {
	if xor.Prev.El == 0 {
		xor.El = el
		xor.Next = &XORLinkedList{}
		xor.Prev = &XORLinkedList{El: -1}
		return
	}

	var newNode *XORLinkedList = &XORLinkedList{}

	for xor.El != 0 {
		// currNode := xor

		if xor.Next.El == 0 {
			newNode.El = el
			newNode.Prev = xor
			xor.Next = newNode
			newNode.Next = &XORLinkedList{}
		}

		xor = xor.Next
	}
}

// Get retrieves an element fro the list based on a index
func (xor *XORLinkedList) Get(idx int) *XORLinkedList {
	counter := 0

	for xor.El != 0 {

		if counter == idx {
			return xor
		}

		counter++
		xor = xor.Next
	}

	return &XORLinkedList{}
}
