package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्ते",
		"こんにちは",
		"привіт",
	}

	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[3:]
	
	fmt.Println(greetings, first, second, third)
}
