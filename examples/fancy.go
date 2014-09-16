package main

import (
	"fmt"
	"github.com/tlorens/keyboard"
)

const ECHO = "StopLookingOverMyShoulder"

func main() {
	var ch int
	var pw [] int

	fmt.Print("Enter Password: ");
	for ch != 10 {
		ch = keyboard.ReadKey()
	    pw = append(pw, ch)
	    fmt.Printf("%c", ECHO[len(pw)-1])
	}

	fmt.Println("\n")
}
