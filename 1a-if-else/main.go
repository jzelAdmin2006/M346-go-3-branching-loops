package main

import "fmt"

func main() {
	maxPoints := 100.0
	scored := 75.0
	ratio := scored / maxPoints

	if ratio > 0.8 {
		fmt.Println("a ratio of", ratio, "is excellent")
	}
	if ratio > 0.6 {
		fmt.Println("a ratio of", ratio, "is good")
	}

	if ratio > 0.8 {
		fmt.Println("a ratio of", ratio, "is excellent")
	} else if ratio > 0.6 {
		fmt.Println("a ratio of", ratio, "is good")
	} else if ratio > 0.4 {
		fmt.Println("a ratio of", ratio, "is acceptable")
	} else {
		fmt.Println("a ratio of", ratio, "is poor")
	}
}
