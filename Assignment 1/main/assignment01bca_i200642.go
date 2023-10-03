package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	assignmentpackage "github.com/naufilmoten/Blockchain/Assignment 1/package"
)

func main() {
	bc := &assignmentpackage.Blockchain{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n\n1. Add a new block")
	fmt.Println("2. View current blocks")
	fmt.Println("3. Verify the chain")
	fmt.Println("4. Change the chain")
	fmt.Println("5. Exit\n\n")

	for {
		fmt.Print("\n\nEnter your choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter transaction data: ")
			scanner.Scan()
			transaction := strings.TrimSpace(scanner.Text())
			fmt.Print("Enter Nonce: ")
			scanner.Scan()
			nonceinput := strings.TrimSpace(scanner.Text())
			nonce, err := strconv.Atoi(nonceinput)
			if err != nil {
				fmt.Println("Invalid Nonce:", err)
				continue
			}
			previousHash := ""
			if len(bc.Blocks) > 0 {
				previousHash = bc.Blocks[len(bc.Blocks)-1].CurrentHash
			}
			block := assignmentpackage.NewBlock(nonce, transaction, previousHash)
			bc.Blocks = append(bc.Blocks, block)
			fmt.Println("Block added to the blockchain.")
		case "2":
			fmt.Println("Blockchain:")
			bc.ListBlocks()
		case "3":
			fmt.Println("Verifying Blockchain.")
			if bc.VerifyChain() {
				fmt.Println("Blockchain is valid.")
			} else {
				fmt.Println("Blockchain is not valid.")
			}
		case "4":
			fmt.Print("Enter block index to change: ")
			scanner.Scan()
			index := strings.TrimSpace(scanner.Text())
			fmt.Print("Enter new transaction data: ")
			scanner.Scan()
			newTransaction := strings.TrimSpace(scanner.Text())
			blockIndex, err := strconv.Atoi(index)
			if err != nil || blockIndex >= 0 && blockIndex < len(bc.Blocks) {
				assignmentpackage.ChangeBlock(bc, blockIndex, newTransaction)
				fmt.Println("Block's transaction changed.")
			} else {
				fmt.Println("Invalid block index.")
			}
		case "5":
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
