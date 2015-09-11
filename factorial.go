package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 30; i++ {
		fmt.Printf("%d: %d\n", i, factorial(i))
	}

}

func factorial(n int) (fac int) {
	if (n == 0) {
		fac = 1
	} else {
		fac = n * factorial(n - 1)
	}
	return 
}
