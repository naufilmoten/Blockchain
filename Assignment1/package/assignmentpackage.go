package assignmentpackage

import (
	"crypto/sha256"
	"fmt"
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

	for i := len(bc.Blocks) - 1; i >= 0; i-- {
		currentBlock := bc.Blocks[i]

		if i == 0 {
			if CalculateHash(currentBlock) != currentBlock.CurrentHash {
				return false
			}
		} else {
			previousBlock := bc.Blocks[i-1]

			if CalculateHash(currentBlock) != currentBlock.CurrentHash {
				return false
			}

			if currentBlock.PreviousHash != previousBlock.CurrentHash {
				return false
			}
		}
	}

	return true
}
