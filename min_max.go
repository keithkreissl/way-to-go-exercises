package main

import (
	"fmt"
)

func main(){

	slice := []int{100, 50, 80, 30, 10, 15 }

	min := minSlice(slice)

	fmt.Println(min)

	max := maxSlice(slice)

	fmt.Println(max)

}


func minSlice(slice []int) (min int){
	for _,v := range slice {
		if 0 == min || v <= min {
			min = v
		}
	}
	return 
}
func maxSlice(slice []int) (max int){
	for _,v := range slice {
		if 0 == max || v >= max {
			max = v
		}
	}
	return 
}
