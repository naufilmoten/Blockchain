package main

import "fmt"

func main() {

	var arr1 = []int{1, 2, 3}
	//creates a variable of a particular type, attaches a name to it, and sets its initial value

	arr2 := [5]int{4, 5, 6, 7, 8} // := is called short variable declaration which takes form
	var arr3 = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr1[0])
	fmt.Println(arr2[2])
	fmt.Println(len(arr2))
}
