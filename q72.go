package main

import (
	"fmt"
)

func main() {
	b := []string{"g","o","l","a","n","g"}
	fmt.Println("o l a", b[1:4])

	fmt.Println("g o", b[:2])

	fmt.Println("l a n g", b[2:])

	fmt.Println("g o l a n g", b[:])
}
