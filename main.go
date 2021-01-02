package main

import (
	"fmt"

	rnd "github.com/tadomikikuto-bit/go-solutions/roman_num_decoder"
	vpwr "github.com/tadomikikuto-bit/go-solutions/valid_palyndrom_with_removal"
	xll "github.com/tadomikikuto-bit/go-solutions/xor_linked_list"
	msa "github.com/tadomikikuto-bit/go-solutions/maxsumarr"
)

func main() {
	dec := rnd.Decode("XXI")
	palyndrom := vpwr.ValidPalyndromWithRemoval("aba")
	xorll := &xll.XORLinkedList{}
	
	maxsum := msa.MaxSumArr([]int{1, 4, 6, 2, 10})
	
	fmt.Printf("%v \n%v \n%v \n%v", dec, palyndrom, maxsum, xorll)
}
