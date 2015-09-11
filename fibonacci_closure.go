package main

import (
	"fmt"
)

func main(){
	value := Fibonacci()	
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d:  %d\n", i, value())
	}
}

func Fibonacci() (func() int) {
	x := 0
	y := 1
	return func () (val int){
		val, x, y = y, y, x+y
		return
	}
}




