package main

import (
	"fmt"
)

func main(){
	PrintlnInts(1,2,3,4,5)

	arr := []int{1,2,3,4,5}

	PrintlnInts(arr...)
}

func PrintlnInts(i ...int){
	for _, v := range i {
		fmt.Println(v)
	}
}

