package main

import (
	"fmt"
)


func main() {
	var	a [5]int

	a[2] = 3
	b := a

	a[3] = 4

	fmt.Println("a[2] ", a[2])
	fmt.Println("b[2] ", b[2])
	fmt.Println("a[3] ", a[3])
	fmt.Println("b[3] ", b[3])
}


