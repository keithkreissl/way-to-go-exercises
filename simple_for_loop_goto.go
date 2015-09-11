package main

import (
	"fmt"
)

func main(){
	i := 0
	S1:
		fmt.Println(i)
		i++
		if i < 15 {
			goto S1
		}
}
