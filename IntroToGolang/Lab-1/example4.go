package main

import "fmt"

func main() {

	elements := []int{999, 99, 9}

	for i := 0; i < len(elements); i++ {
		fmt.Print(elements[i] + 1)
		fmt.Print(" ")
	}
	fmt.Println("....DONE!")
}

// Example 4 Done. Allah Hafiz
