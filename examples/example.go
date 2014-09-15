package main

import (
	"fmt"
	"github.com/tlorens/keyboard"
)

func main() {
	var ch int

	fmt.Println("Press END to quit");
	for ch != 152 {
	    // Good for special keys: Arrows, PgUp, PgDn, Home, End.
	    ch = keyboard.ReadKey()
	    fmt.Printf("[%d][%c]\n", ch, ch);   
	}
}