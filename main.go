package main

import (
	"fmt"
	"github.com/tamsinforbes/tjf-go-notes/emoji_math"
)

func main() {

	fmt.Println("Yo!")

	fmt.Printf("%c + %c = %c\n", 'a', 'b', emoji_math.emoji_add('a', 'b'))
}