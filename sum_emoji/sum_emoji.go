package main

import (
  "fmt"
	"math"
)

// Declare variable names and type for use in subsequent functions
var (
	x int
	y int
	a rune
	b rune
)

// Write smaller functions in the same script as the main 
// function that uses them; cannot import from local directories
// only public sources

func add_int(x int, y int) int {
	return x + y 
	// this x & y are the arguments declared in this function
	// not the variables declaed above
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
// Shorten add_rune(x rune, y rune) to add_rune(x, y rune)

func add_rune(x, y rune) rune {
	return x + y
}

// A 'naked' return only returns the named values, 
// named in the function first line
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func get_integer_part_of_decimal(decimal float64) (x float64, y int) {
	x = decimal
	y = int(decimal)
	return
}

func print_floor_decimal(decimal float64) (x float64, y int) {
	x = decimal
	y = int(decimal) // Gets the integer part, or floor, does not round up
	fmt.Printf("Using int() returns the floor of the decimal: %v -> %v \n", x, y)
	return
}

func print_round_decimal(decimal float64) (x float64, y int) {
	x = decimal
	y = int(math.Round(decimal)) // Round, then convert to int
	fmt.Printf("Round and then get integer part: int(math.Round(decimal)): %v -> %v\n", x, y)
	return
}

// Declare variables outside a function
// If all are the same type, declare once after the final variable
// Declared variables are given a default value
var c, python, java bool


// main function takes no variables? just defines what other
// functions to run 
func main() {
	x = 13
  y = 42
	int_sum := add_int(x, y)
	fmt.Printf("%d + %d = %d \n", x, y, int_sum)

	a = 'a'
	b = 'b'
	rune_sum := add_rune(a, b)
	fmt.Printf("%c + %c = %c \n", a, b, rune_sum)
	fmt.Printf("(%U + %U = %U) \n", a, b,  rune_sum)

	fmt.Println(split(17))
	
	fmt.Println(get_integer_part_of_decimal(1.234))

	print_floor_decimal(2.6345)

	print_round_decimal(7.89)

	// Declare variables inside a function
	var i, j int
	// Println prints to a single line and separates operands with spaces
	// Formatting is set to its defaults
	fmt.Println("Default values for int and bool variables:", i, j, c, python, java)

	// Assign values to the variables
	i, j, c, python, java = 2, 3, true, true, false
	fmt.Println("Default values for int and bool variables:", i, j, c, python, java)

}
