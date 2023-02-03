package main

import (
	"fmt"
	"github.com/sebastianterleira/Hello-paquetes-GO/greet"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(greet.Spanish())
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())
}
