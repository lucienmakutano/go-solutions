package main

import (
	"fmt"

	rnd "github.com/tadomikikuto-bit/go-solutions/roman_num_decoder"
	vpwr "github.com/tadomikikuto-bit/go-solutions/valid_palyndrom_with_removal"
)

func main() {
	dec := rnd.Decode("XXI")
	palyndrom := vpwr.ValidPalyndromWithRemoval("aba")
	fmt.Printf("%v \n%v", dec, palyndrom)
}
