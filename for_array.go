package main

import (
	"fmt"
)

func main() {
	var a [16]int

	for i := range a {
		a[i] = i
	}

	for i := range a {
		fmt.Println("i: ", i, " value: ", a[i]);
	}
}


