package main

import (
	"fmt"
	"github.com/tlorens/keyboard"
)

func main() {
	var ch int

	fmt.Println("Press TAB to quit");
	for ch != 9 {
	    // Good for special keys: Arrows, PgUp, PgDn, Home, End.
	    ch = keyboard.RawKey()
	    fmt.Printf("[%d][%c]\n", ch, ch);   
	}
}