package main

import (
	"fmt"
)

const (
	FIB_SIZE=50
)

func main(){
	f := fib(FIB_SIZE)
	for i := 0; i < FIB_SIZE; i++{
		fmt.Println(f())
	}

}
func fib(size int)(f func() int) {
	index := 0
	a := make([]int, size)
	f =  func () (value int) {
		if (index - 2 < 0) {
			value = 1
		} else {
			value = a[index - 2] + a[index - 1]
		}
		a[index] = value
		index++
		return 
	}
	return 
}	
