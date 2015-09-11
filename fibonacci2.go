package main

import (
	"fmt"
)

func main(){
	for i := 0; i <= 10; i++ {
		value, pos := Fibonacci(i)	
		fmt.Printf("%d and %d\n", value, pos)
	}
}

func Fibonacci(num int) (value int, pos int){
	pos = num
	if (num <= 1) {
		value = 1
	} else {
		val1,_ := Fibonacci(num - 2)
		val2,_ := Fibonacci(num -1)
		value = val1 + val2
	}
	return
}


