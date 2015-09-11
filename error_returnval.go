package main

import (
	"fmt"
	"math"
	"errors"
)

func main(){
	var num1 float64 = 64
	var num2 float64 = -64

	PrintSqrtUnnamed(num1)
	PrintSqrtUnnamed(num2)
	PrintSqrtNamed(num1)
	PrintSqrtNamed(num2)

}

func PrintSqrtUnnamed(num float64){
	value, err := MySqrtUnnamed(num)
	if (err != nil){
		fmt.Printf("Error for number %f: %s\n", num, err)
	} else {
		fmt.Printf("The square root of %f is %f\n", num, value)
	}
}
func PrintSqrtNamed(num float64){
	value, err := MySqrtNamed(num)
	if (err != nil){
		fmt.Printf("Named: Error for number %f: %s\n", num, err)
	} else {
		fmt.Printf("Named: The square root of %f is %f\n", num, value)
	}
}

func MySqrtUnnamed(num float64)(float64, error){
	if num < 0 {
		return 0, errors.New("number must be greater than 0")
	}
	return math.Sqrt(num), nil
}

func MySqrtNamed(num float64)(val float64, err error){
	if num < 0{
		err = errors.New("number must be greater than 0")
		return
	}
	val = math.Sqrt(num)
	return
}



