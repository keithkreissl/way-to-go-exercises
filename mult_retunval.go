package main

import (
	"fmt"
)

func main(){
	a, b, c := addMultiplyDiffUnnamed(4,2)
	printValues("unnamed", a, b, c)

	a, b, c = addMultiplyDiff(4, 2)
	printValues("named", a, b, c)

}

func addMultiplyDiffUnnamed(a int, b int)(int, int, int){
	return a + b, a * b, a - b
}

func addMultiplyDiff(a int, b int)(add int, mult int, diff int){
	add = a + b
	mult = a * b
	diff = a - b
	return
}

func printValues(name string, a int, b int, c int){
	fmt.Printf("name: %s\n Add: %d\n Multiply: %d\n Diff: %d\n", name, a, b, c)	
}
