package main

import (
	"fmt"
	"github.com/tlorens/keyboard"
)

func main() {
	var ch int

	fmt.Print("Do you wish to continue? (Y/n): ")

    ch = keyboard.ReadKey()

    switch(ch) {
    	case 89,121:
    		fmt.Print("Yes")
    	case 78,110:
    		fmt.Print("No")
    	default:
			fmt.Print("Yes")
    }

    fmt.Println("\n\n")
}
