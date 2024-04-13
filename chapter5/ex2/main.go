package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("Cannot open a file")
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Cannot get a file stat")
	}

	fmt.Println(fileInfo.Size())
}
