package publicmodule

import (
	"fmt"

	"rsc.io/quote"
)

func Hello(name string) {
	fmt.Println("Hello ", name)
}

func GetQuote() {
	q := quote.Go()
	fmt.Println(q)
}
