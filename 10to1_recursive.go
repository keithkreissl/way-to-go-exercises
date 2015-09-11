package main

import (
	"fmt"
)

func main() {
	printrec(10)
}

func printrec(i int){
	fmt.Println(i)
	if ( i > 1 ) { 
		printrec(i - 1)
	}
}
