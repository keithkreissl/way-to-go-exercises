package main

import (
	"fmt"
)

func main() {
	fn := func(s string) {
		fmt.Println(s)
	}

	f(fn);
}

func f(a func(string)){
	a("Hello World")
}
