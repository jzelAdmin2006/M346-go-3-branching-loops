package main

import (
	"fmt"
	"strconv"
)

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i <= Upper; i++ {
		var result string
		switch i {
		case 3:
			fallthrough
		case 6:
			fallthrough
		case 9:
			fallthrough
		case 12:
			fallthrough
		case 18:
			fallthrough
		case 21:
			fallthrough
		case 24:
			fallthrough
		case 27:
			result = "Fizz"
		case 5:
			fallthrough
		case 10:
			fallthrough
		case 20:
			fallthrough
		case 25:
			result = "Buzz"
		case 15:
			fallthrough
		case 30:
			result = "FizzBuzz"
		default:
			result = strconv.Itoa(i)
		}
		fmt.Println(result)
	}
}
