package main 

import (
	"fmt"
)

func main() {
	s := make([]byte, 5)
	fmt.Println("length: 5 ", len(s))
	fmt.Println("cap: 5 ", cap(s))

	s = s[2:4]
	fmt.Println("length: 2 ", len(s))
	fmt.Println("cap: 3 ", cap(s))

}
