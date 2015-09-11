package main

import (
	"fmt"
)

func main() {
//its a slice because you are not giving a size
//you send a copy of the SliceHeader but it operates on the underlying data array
	h := []byte{'H', 'e', 'l', 'l', 'o'}
	w := []byte{'w', 'o', 'r', 'l', 'd'}

	hw := Append(h, w)
	fmt.Println(string(hw))

	foo := make([]byte, 3, 25)
	foo[0] = 'f'
	foo[1] = 'o'
	foo[2] = 'o'

	bar := []byte{'b', 'a', 'r'}

	foobar := Append(foo, bar)
	fmt.Printf("%s, %d\n", string(foobar), len(foobar))
}

func Append(sl []byte, data []byte)  []byte{
	index := len(sl) 
	if (cap(sl) < len(sl) + len(data)) {
		tmp := make([]byte, len(sl) + len(data))
		for i,v := range sl {
			tmp[i] = v
		}
		sl = tmp
	} else {
		sl = sl[0:len(sl) + len(data)]
	}
	for i,v := range data {
		sl[index + i] = v
	}
	return sl
}
