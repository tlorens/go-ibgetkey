package main

import (
	"fmt"
	"github.com/tlorens/keyboard"
)

func main() {
	var ch int

	fmt.Print("Enter Password: ");
	for ch != 10 {
	    ch = keyboard.ReadKey()
	    fmt.Print("*");
	}
}
