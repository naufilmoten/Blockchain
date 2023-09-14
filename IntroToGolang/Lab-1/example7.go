package main

import "fmt"

func main() {
	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence
	fmt.Printf("strArrayl: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	strArray1[0] = "Canada"
	fmt.Printf("strArrayl: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %vi\n", *strArray3)
}
