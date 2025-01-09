package main

import (
  "fmt"
  "strconv"
	"io"
	"os"
)

var (
	a int
	b int
	mynum int
)

func main() {
  fmt.Println("INTEGER LITERALS:")
	a = 5
	b = 2
	mynum = a + b
	mystring := strconv.Itoa(mynum) // Itoa = Int to ASCII
	fmt.Printf("%s %s\n", mystring, strconv.FormatInt(int64(a), 10))
	fmt.Printf("%s\n", strconv.Itoa(mynum))
	fmt.Printf("%s\n", strconv.FormatInt(int64(mynum), 10))
	fmt.Printf("Using an underscore for readability does not change the integer: 4_2 = %s\n", strconv.FormatInt(int64(4_2), 10))
	fmt.Printf("4_2_1 = %s\n", strconv.FormatInt(int64(4_2_1), 10))
	fmt.Printf("42 = %s\n", strconv.FormatInt(int64(42), 10))
	fmt.Printf("Converting octal 0600 to string with base 8 prints octal representation of 384 (6x64): 0600 = %s\n", strconv.FormatInt(int64(0600), 8))
	fmt.Printf("Converting octal 0600 to string with base 10 prints decimal representation of 384: 0600 = %s\n", strconv.FormatInt(int64(0600), 10))
	fmt.Printf("Converting octal 0600 to string with base 2 prints: 0600 = %s\n", strconv.FormatInt(int64(0600), 2))
	fmt.Printf("Converting binary 0b110000000 to string with base 10 prints decimal representation of 384 (256+128): 0b110000000 = %s\n", strconv.FormatInt(int64(0b110000000), 10))
	fmt.Printf("If you use the prefixes 0b, 0o, 0x for binary, octal and hex numbers respectively\nthen when you hover over them in VSCode their decimal equivalents appear.\n")
	fmt.Printf("For Hex the Unicode/UTF-8 character the number represents also appears! Eg 0x2c (U+002C) is comma ','.\n")
	c := 0b101100
	fmt.Printf("Binary 0b101100 is %s in decimal\n", strconv.FormatInt(int64(c), 10))
	d := 0o54
	fmt.Printf("Octal 0o54 is %s in decimal\n", strconv.FormatInt(int64(d), 10))
	e := 0x2c
	fmt.Printf("Hex 0x2c is %s in decimal\n", strconv.FormatInt(int64(e), 10))
	f := 0xbd
	fmt.Printf("Hex 0xbd is %s (11x16+13) in decimal and represents UTF-8 U00BD: \u00bd\n", strconv.FormatInt(int64(f), 10))

	fmt.Println("\nFLOATING POINT LITERALS")
	// g := 0x1.Fp+0 
	// fmt.Printf("%s\n", strconv.FormatFloat(float64(g), "b", -1, ))
	fmt.Printf("!!!USE the fmt package string format codes!!!\nThen all the string conversion rubbish is not required.\n")
	fmt.Printf("Base 2 format code '%%b' (b for binary) prints 0b110 as: %b \n", 0b110)
	fmt.Printf("Base 10 format code '%%d' (d for decimal) prints 44 as: %d \n", 44)
	fmt.Printf("Base 2 format code '%%b' converts decimal 44 to binary and prints string: %b \n", 44)
	fmt.Printf("Base 10 format code '%%d' converts binary 0b110 to decimal and prints string: %d \n", 0b110)
	fmt.Printf("Base 8 format code '%%o' (o for octal) prints 0o250 as: %o \n", 0o250)
	fmt.Printf("Base 10 format code '%%d' converts octal 0o250 to decimal: %d \n", 0o250)
	fmt.Printf("Base 2 format code '%%b' converts octal 0o250 to binary: %b \n", 0o250)
	fmt.Printf("Base 16 format code '%%x' (x for hex) prints 0x1a as: %x \n", 0x1a)
	fmt.Printf("Base 10 format code '%%d' converts hex 0x1a to decimal: %d \n", 0x1a)
	fmt.Printf("Base 2 format code '%%b' converts hex 0x1a to binary: %b \n", 0x1a)
	fmt.Printf("Base 8 format code '%%o' converts hex 0x1a to octal: %o \n", 0x1a)
	fmt.Printf("Unicode format code '%%U' prints the unicode code point (not the value) for the equivalent hex: %x = %U \n", 0x1a, 0x1a)
	fmt.Printf("Unicode format code '%%U' prints the unicode for the equivalent hex: %x = %U \n", 0xbd, 0xbd)
	fmt.Printf("Unicode format code '%%c' prints the unicode value for a unicode code point in hex: %c = %U \n", 0x1D70B, 0x1D70B)
	fmt.Printf("To print the actual value of a Unicode code just escape the code and type in the string directly: \\u1234 == \u1234\n")
	fmt.Printf("Unicode code as typed: \\u00bd == \u00bd\n")
	fmt.Printf("Unicode properly formatted: %U == \u00bd\n", 0x00bd)

	fmt.Printf("\nString interpolation methods\n")
	const name, age = "Kim", 42
	s := fmt.Sprintln(name, "is", age, "years old.")
	io.WriteString(os.Stdout, s) // Ignoring error for simplicity. What error?
	s = fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s) // Ignoring error for simplicity. What error?

	fmt.Printf("\nIMAGINERY LITERALS\n")
	// Complex numbers format as parenthesized pairs of floats, with an 'i'
	// after the imaginary part.
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.5f %.2e\n", point, point, point, point, point)
	// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
 
	fmt.Printf("\nRUNE LITERALS\n")
	// Runes are integers but when printed with %c show the character with that
	// Unicode value. The %q verb shows them as quoted characters, %U as a
	// hex Unicode code point, and %#U as both a code point and a quoted
	// printable form if the rune is printable. %v is for numerical value in decimal.
	smile := '😀'
	fmt.Printf("%%v: %v, %%d: %d, %%c: %c, %%q: %q, %%U: %U, %%#U: %#U, %%c with 0x1F600: %c \n", smile, smile, smile, smile, smile, smile, 0x1F600)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	fmt.Printf("Addition with Runes: %c + %c = %c", smile, smile, smile + smile)
	fmt.Printf("Addition with Runes: %c + %c = %c, %v \n", smile, smile, smile + smile, smile + smile)
	lower_d := 'd'
	upper_e := 'E'
	fmt.Printf("Addition with Runes: %c + %c = %c; %v + %v = %v \n", lower_d, upper_e, lower_d + upper_e, lower_d, upper_e, lower_d + upper_e)
	hedgehog := '🦔'
	fmt.Printf("Addition with Runes: %c + 'h' = %c \n", hedgehog, hedgehog+'h')

	heart := '♥'
	fmt.Printf("Addition with Runes: %c + %c = %c \n", heart, heart, heart+heart)

	sun := '☼'
	//cloud := '☁️'
	// suncloud := '🌤️'
	fmt.Printf("Addition with Runes: %c + %c = %c,  \n", sun, lower_d, sun + lower_d)

	fmt.Printf("Sum of cat: %c + %c + %c = %c \n", 'c', 'a', 't', 'c' + 'a' + 't')
	fmt.Printf("Sum of z+z+z: %c + %c + %c = %c (%v, %U) \n", 'z', 'z', 'z', 'z'+'z'+'z', 'z'+'z'+'z', 'z'+'z'+'z')
	fmt.Printf("3 times z: 3 * %c = %c (%v, %U) \n", 'z', 3*'z', 3*'z', 3*'z')
	fmt.Printf("%c * %c = %c (%v, %U) \n",'a', 'z', 'a'*'z', 'a'*'z', 'a'*'z')
	Jon := 'J'+'o'+'n'
	pi := '𝜋'
	fmt.Printf("Jon: %c + %c + %c = %c Planck's constant over 2%c (%v, %U) \n", 'J', 'o', 'n', Jon, pi, Jon, Jon)

	circle_one := '①'
	circle_two := '②'
	fmt.Printf("%c + %c = %c (%v, %U) \n", circle_one, circle_two, circle_one+circle_two, circle_one+circle_two, circle_one+circle_two)

	leopard := '🐆'
	cat := '🐈'
	fmt.Printf("%c + 2 = %c (%v %U %d) \n", leopard, leopard + 2, cat, cat, cat)
	tam := 't'+'a'+'m'
	fmt.Printf("tam: %c (%v, %U) \n", tam, tam, tam)
	tamsin := 't'+'a'+'m'+'s'+'i'+'n'
	fmt.Printf("tamsin: %c (%v, %U) \n", tamsin, tamsin, tamsin)

	// To format a Unicode character direct from the hex value use %c
	fmt.Printf("Use %%c to format a hex number as its Unicode character: 0x1D70B = %c\n", 0x1D70B)
	fmt.Printf("Use %%c to format a hex number as its Unicode character: 0x0F3A = %c\n", 0x0F3A)
	fmt.Printf("Use %%v to format a hex number as its decimal value: 0x0F3A = %v\n", 0x0F3A)

	fmt.Printf("\nSTRING LITERALS\n")
	// Single quotes are reserved for runes so cannot be used interchangeably with double quotes
	fmt.Println(`Use double quotes for interpretated strings:  ticks for raw strings: abc\n`)
	fmt.Println(`Use back ticks for raw strings: abc\n`)
	fmt.Println(`Raw strings between back ticks can span 
	multiple lines in code, but prints those 
		new lines too and indentation is 
		not as expected`)
	fmt.Println(`
		This
		works to
		align text to itself,
		but includes the extra newline
		at beginning and end.
	`)



}