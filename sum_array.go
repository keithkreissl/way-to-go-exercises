package main

import (
	"fmt"
)

func main(){
	arrF := []float32{10, 20, 30, 40}

	sum := Sum(arrF)
	fmt.Println(sum)
}

func Sum(arrF []float32) (sum float32){
	for _,v := range arrF {
		sum = sum + v
	}
	return 
}
