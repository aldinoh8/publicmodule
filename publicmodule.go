package publicmodule

import (
	"fmt"

	"rsc.io/quote"
)

func Hello(name string) {
	fmt.Println("Hello ", name, "from v0.2.0")
}

func GetQuote() {
	q := quote.Go()
	fmt.Println(q)
}
