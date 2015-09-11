package main

import "fmt"

func main(){
	for i := 1; i < 26; i++ {
		fmt.Printf("line %d:", i)
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
}
