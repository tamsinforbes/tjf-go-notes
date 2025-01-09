package main

import (
	"fmt"
)

// Function for use by the executable entry point
// functions_in-multiple_files/main.go 
 
func print_decimal_in_other_bases(x int) {
	fmt.Printf("Decimal: %d, Hex: %x, Octal: %o, Binary: %b\n", x, x, x, x)
}