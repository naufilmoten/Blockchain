package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Haris", 10000, "UI/UX Designer"}
	emp3 := employee{"Jack", 90000, "Backend Developer"}

	employees := []employee{emp1, emp2, emp3}

	companyDetails := company{"Tetra", employees}

	fmt.Println("Company Name:", companyDetails.companyName)

	fmt.Println("Employee 1:")
	fmt.Println("Name:", companyDetails.employees[0].name)
	fmt.Println("Salary:", companyDetails.employees[0].salary)
	fmt.Println("Position:", companyDetails.employees[0].position)

	fmt.Println("Employee 2:")
	fmt.Println("Name:", companyDetails.employees[1].name)
	fmt.Println("Salary:", companyDetails.employees[1].salary)
	fmt.Println("Position:", companyDetails.employees[1].position)

	fmt.Println("Employee 3:")
	fmt.Println("Name:", companyDetails.employees[2].name)
	fmt.Println("Salary:", companyDetails.employees[2].salary)
	fmt.Println("Position:", companyDetails.employees[2].position)
}
