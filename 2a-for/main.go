package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	j := 0
	for j < 10 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Printf("%d ", x)
	}
	fmt.Println()

	sum, maxSum := 0, 15
	for y := 0; y < 10; y++ {
		sum += y
		fmt.Printf("%d", y)
		if sum >= maxSum {
			break
		} else {
			fmt.Print(" + ")
		}
	}
	fmt.Println(" >=", maxSum)
}
