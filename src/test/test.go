package main

import (
	"fmt"
)

func main() {
	a := 1
	b := (a | (8))

	//00001000
	fmt.Println("b -->", b)
}
