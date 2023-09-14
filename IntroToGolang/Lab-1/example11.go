package main

import (
	"crypto/sha256"
	"fmt"
)

func CalculateHash(stringToHash string) string {
	fmt.Printf("String Received: %s\n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}
func main() {

	fmt.Printf("%x\n", CalculateHash("Naufil"))
}
