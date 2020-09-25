package main

import (
	"fmt"
	"github.com/tadomikikuto-bit/go-solutions/roman_num_decoder"
)

func main(){
	dec := roman_num_decoder.Decode("XXI")
	fmt.Println(dec)
}
