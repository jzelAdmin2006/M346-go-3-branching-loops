package main

import "fmt"

func main() {
	grade := 5

	if grade == 6 {
		fmt.Println("very good")
	} else if grade == 5 {
		fmt.Println("good")
	} else if grade == 4 {
		fmt.Println("sufficient")
	} else if grade == 3 {
		fmt.Println("insufficient")
	} else if grade == 2 {
		fmt.Println("bad")
	} else if grade == 1 {
		fmt.Println("very bad")
	}

	switch grade {
	case 6:
		fmt.Println("very good")
	case 5:
		fmt.Println("good")
	case 4:
		fmt.Println("sufficient")
	case 3:
		fmt.Println("insufficient")
	case 2:
		fmt.Println("bad")
	case 1:
		fmt.Println("very bad")
	default:
		fmt.Println("unknown grade")
	}

	switch grade {
	case 6:
		fallthrough
	case 5:
		fallthrough
	case 4:
		fmt.Println("passed the exam")
	case 3:
		fallthrough
	case 2:
		fallthrough
	case 1:
		fmt.Println("failed the exam")
	default:
		fmt.Println("unknown result")
	}
}
