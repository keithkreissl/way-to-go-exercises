package main

import (
	"fmt"
	"bytes"
)

func main() {

	getNextString := nextString()

	var buffer bytes.Buffer

	for {
		if s, ok := getNextString(); ok{
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")

}

func nextString() (func() (string, bool)) {
	i := 1
	return func() (value string, ok bool){
		value = fmt.Sprintf("%d ", i)
		if i % 100 == 0 {
			ok = false
		} else {
			ok = true
		}
		i++
		return 
	}
}




