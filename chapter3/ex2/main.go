package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"

	runes_slice := []rune(message)

	fmt.Printf("rune is %c", runes_slice[3])
}
