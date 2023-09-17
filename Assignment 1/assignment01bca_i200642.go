package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

type Blockchain struct {
	Blocks []*Block
}

func CalculateHash(b *Block) string {
	combinedstring := fmt.Sprintf("%d%s%s", b.Nonce, b.Transaction, b.PreviousHash)
	hash := sha256.Sum256([]byte(combinedstring))
	return fmt.Sprintf("%x", hash)
}

func NewBlock(nonce int, transaction string, previousHash string) *Block {
	block := new(Block)

	block.Nonce = nonce
	block.Transaction = transaction
	block.PreviousHash = previousHash

	block.CurrentHash = CalculateHash(block)

	return block
}

func (bc *Blockchain) ListBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println()
	}
}

func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bc.Blocks) {
		block := bc.Blocks[index]
		block.Transaction = newTransaction

	}
}

func (bc *Blockchain) VerifyChain() bool {
	if len(bc.Blocks) == 0 {
		return true
	}

	for i := len(bc.Blocks) - 1; i > 0; i-- {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if CalculateHash(currentBlock) != currentBlock.CurrentHash {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}

	return true
}

func main() {
	bc := &Blockchain{}
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
			block := NewBlock(nonce, transaction, previousHash)
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
				bc.ChangeBlock(blockIndex, newTransaction)
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

// 	bc.Blocks = append(bc.Blocks, NewBlock(0, "Genesis Transaction", ""))
// 	bc.Blocks = append(bc.Blocks, NewBlock(1, "Alice to Bob", bc.Blocks[len(bc.Blocks)-1].CurrentHash))
// 	bc.Blocks = append(bc.Blocks, NewBlock(2, "Bob to Carol", bc.Blocks[len(bc.Blocks)-1].CurrentHash))

// 	// Print the blockchain
// 	fmt.Println("Blockchain:")
// 	bc.ListBlocks()

// 	// Change a block's transaction
// 	fmt.Println("Changing block 1...")
// 	bc.ChangeBlock(1, "Mallory to Eve")

// 	// Print the blockchain after change
// 	fmt.Println("Blockchain after change:")
// 	bc.ListBlocks()

// 	// Verify blockchain integrity
// 	fmt.Println("Verifying blockchain integrity...")
// 	if bc.VerifyChain() {
// 		fmt.Println("Blockchain is valid.")
// 	} else {
// 		fmt.Println("Blockchain is not valid.")
// 	}
// }
