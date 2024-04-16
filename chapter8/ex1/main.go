package main

import "fmt"

type Num interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Double[T Num](a T) T {
	return 2 * a
}

func main(){
	fmt.Println(Double(5))
	fmt.Println(Double(3.23))
}
