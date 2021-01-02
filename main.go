package main

import (
	"fmt"

	rnd "github.com/tadomikikuto-bit/go-solutions/roman_num_decoder"
	vpwr "github.com/tadomikikuto-bit/go-solutions/valid_palyndrom_with_removal"
	xll "github.com/tadomikikuto-bit/go-solutions/xor_linked_list"
)

func main() {
	dec := rnd.Decode("XXI")
	palyndrom := vpwr.ValidPalyndromWithRemoval("aba")
	var xorll *xll.XORLinkedList = &xll.XORLinkedList{}

	fmt.Printf("%v \n%v \n%v", dec, palyndrom, xorll)
}
