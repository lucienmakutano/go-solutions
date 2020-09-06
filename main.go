package main

import (
	"fmt"
	"github.com/tadomikikuto-bit/roman_num_dec"
)

func main(){
	dec := roman_num_dec.Decode("XXI")
	fmt.Println(dec)
}
