package main

import (
	"fmt"
	"github.com/sebastianterleira/Hello-paquetes-GO/greet"
	"rsc.io/quote"
)

func main() {
	fmt.Println(greet.Spanish())
	fmt.Println(quote.Hello())
}
