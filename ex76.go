package main

import (
	"fmt"
	"bytes"
)

func main() {

	h := "Hello"

	w := " World!"
	hw := Append([]byte(h), []byte(w))
	fmt.Println(string(hw))
}

func Append(sl []byte, data []byte) (ret []byte){
	buffer := bytes.NewBuffer(sl)

	if len(data) + len(sl) > cap(sl){
		buffer.Grow(len(data) + len(sl))
	}

	buffer.Write(data)

	ret = buffer.Bytes()

	return
}
