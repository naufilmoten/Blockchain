package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	Subjects   []string
}

func NewStudent(rollno int, name string, address string, Subs []string) *Student {

	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.Subjects = Subs
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, Subs []string) *Student {
	st := NewStudent(rollno, name, address, Subs)
	ls.list = append(ls.list, st)

	return st
}

func (s *Student) CalculateBlockHash() string {
	blockData := fmt.Sprintf("%d%s%s%v", s.rollnumber, s.name, s.address, s.Subjects)
	return CalculateHash(blockData)
}

func CalculateHash(stringToHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func (s *StudentList) PrintStudents() {
	for i := 0; i < len(s.list); i++ {
		student := s.list[i]
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Roll Number: %d\n", student.rollnumber)
		fmt.Printf("Name: %s\n", student.name)
		fmt.Printf("Address: %s\n", student.address)
		fmt.Printf("Block Hash: %v\n", student.CalculateBlockHash())
	}
}

func main() {

	Students := new(StudentList)
	arr1 := []string{"AI", "SE", "DB", "OS", "Algo"}
	Students.CreateStudent(1, "John", "Islamabad", arr1)
	arr2 := []string{"DevOps", "Blockchain", "PDC", "PPIT", "Info Sec"}

	Students.CreateStudent(2, "Ali", "Karachi", arr2)
	Students.PrintStudents()

}
