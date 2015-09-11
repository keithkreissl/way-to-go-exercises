package main

import ( 
	"fmt"
	"strings"
)

func main(){
	s := "G"
	for i := 1; i < 26; i++ {
		fmt.Println(strings.Repeat(s,i))
	}
}
