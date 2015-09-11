package main

import (
	"fmt"
)

const (
	FIB_SIZE=50
)

func main(){
	populate, slice := fib(FIB_SIZE)
	populate()

	fmt.Println(slice(10))	

	

}
func fib(size int)(populator func(), slicer func(end int) []int) {
	a := make([]int, size)
	populator =  func () {
		var value int
		for i := 0; i < size; i++{
			if (i - 2 < 0) {
				value = 1
			} else {
				value = a[i - 2] + a[i - 1]
			}
			a[i] = value
		}
	}
	slicer = func(end int) (value []int){
		value = a[:end]
		return 
	}
	return 
}	
