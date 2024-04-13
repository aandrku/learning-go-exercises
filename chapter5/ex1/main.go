package main

import (
	"errors"
	"fmt"
) 

var (
	add = func (i, j int) int { return i + j}
	sun = func (i, j int) int { return i - j}
	div = func (i, j int) (int, error ) { 
		if j == 0 {
			return 0, errors.New("division by zero")
		}
		return i / j, nil
	}
	mul = func (i, j int) int { return i * j}
)

func main() {
	fmt.Println(div(5, 0))
}


//ain't making a calculator app
