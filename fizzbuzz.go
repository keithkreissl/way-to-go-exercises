package main

import "fmt"

func main(){

	for i := 1; i < 101; i++{
		fmt.Printf("%d: ", i)
		switch  {
			case i % 3 == 0 && i % 5 == 0:
				fmt.Print("FizzBuzz")
			case i % 3 == 0:
				fmt.Print("Fizz")
			case i % 5 == 0:
				fmt.Print("Buzz")
		}
		fmt.Println()
	}

}
