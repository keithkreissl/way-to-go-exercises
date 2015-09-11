package main

import "fmt"

func main(){

	str := "Hello, WOrld"

	for pos, char := range str {
		fmt.Printf("%d: %c\n", pos, char)
	}

}
