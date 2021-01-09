package main

import (
	"fmt"

	msa "github.com/tadomikikuto-bit/go-solutions/maxsumarr"
	vpwr "github.com/tadomikikuto-bit/go-solutions/palyndrom"
	rnd "github.com/tadomikikuto-bit/go-solutions/roman_num_decoder"
	xll "github.com/tadomikikuto-bit/go-solutions/xorlinkedlist"
)

func main() {
	dec := rnd.Decode("XXI")
	palyndrom := vpwr.ValidPalyndromWithRemoval("aba")
	xorll := xll.New()
	for i := 0; i < 10; i++ {
		xorll.Insert(i)
	}

	maxsum := msa.MaxSumArr([]int{1, 4, 6, 2, 10})

	fmt.Printf("%v \n%v \n%v\n", dec, palyndrom, maxsum)
	xll.PrintList(xorll)
}
