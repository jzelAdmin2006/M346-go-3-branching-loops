package main

import "fmt"

func main() {
	fibs := []int{1, 1, 2, 3, 5, 8}
	for i := 0; i < len(fibs); i++ {
		value := fibs[i]
		fmt.Printf("%d: %d\n", i, value)
	}
	for i, value := range fibs {
		fmt.Printf("%d: %d\n", i, value)
	}
	for i := range fibs {
		value := fibs[i]
		fmt.Printf("%d: %d\n", i, value)
	}
	for _, value := range fibs {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	zipCodes := map[int]string{
		1200: "Geneva",
		3000: "Bern",
		6000: "Lucerne",
		7000: "Chur",
		8000: "Zurich",
	}
	for zipCode, town := range zipCodes {
		fmt.Printf("%d %s\n", zipCode, town)
	}
}
