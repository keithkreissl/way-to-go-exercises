package main

import (
	"fmt"
)

const (
	FIB_SIZE=50
)

func main(){
	a := make([]int, FIB_SIZE)
	populate(a)

	slice := slicer(a)
	
	fmt.Println(slice(11))
	fmt.Println(slice(5))
	fmt.Println(slice(20))
}



func fibinacci(a []int, index int) (value int){
	if 	(index - 2 < 0) {
		value = 1
	} else {
		value = a[index - 2] + a[index - 1]
	}
	a[index] = value
	return
}

func populate(a []int) {
	for i := 0; i < len(a); i++ {
		fibinacci(a, i)
	}
}

func slicer(a []int) (slice func(end int) []int){
	slice = func(end int) (s []int){
		s = a[:end]
		return
	}
	return
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
