package main

import (
	"fmt"
	"github.com/icebrkr/ibkey"
)

func main() {
	var ch int

	ch = ibkey.ReadKey()
	fmt.Printf("[%d]",ch)
}
