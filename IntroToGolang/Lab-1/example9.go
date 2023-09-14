package main

import "fmt"

type Doctor struct {
	number     int
	doctorName string
	patients   []string
}

func main() {

	aDoctor := Doctor{
		number:     006,
		doctorName: "Naveed",
		patients:   []string{"A", "B", "C", "D"},
	}

	fmt.Printf("Doctor ID: %v, Doctor Name: %v, Patients: %v", aDoctor.number, aDoctor.doctorName, aDoctor.patients)
}
