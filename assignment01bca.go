package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	Index       int
	Transaction string
	Nonce       int
	Prev_Hash   string
	Cur_Hash    string
}

func NewBlock(index int, transaction string, nonce int, previousHash string) *Block {
	//A method to add new block. To keep things simple, you could provide a sting of your choice as a transaction (e.g., “bob to alice”). Also, use any integer value as a nonce. The CreateHash() method will provide you the Block Hash value.
	b := &Block{
		Index:       index,
		Transaction: transaction,
		Nonce:       nonce,
		Prev_Hash:   previousHash,
	}
	//function call based on given protottype
	b.Cur_Hash = CalculateHash(fmt.Sprintf("%d%s%d%s%s", b.Index, b.Transaction, b.Nonce, b.Prev_Hash, b.Cur_Hash))
	return b
}

func ListBlocks(blocks []*Block) {
	//A method to print all the blocks in a nice format showing block data such as transaction, nonce, previous hash, current block hash
	ln := strings.Repeat("=", 55)

	for b := range blocks {
		fmt.Println(ln)
		fmt.Printf("|| Index:       %d\n", blocks[b].Index)
		fmt.Printf("|| Transaction: %s\n", blocks[b].Transaction)
		fmt.Printf("|| Nonce:       %d\n", blocks[b].Nonce)
		fmt.Printf("|| Prev_Hash:   %s\n", blocks[b].Prev_Hash)
		fmt.Printf("|| Cur_Hash:    %s\n", blocks[b].Cur_Hash)
		fmt.Println(ln)
	}

}

func ChangeBlock(b *Block, newTransaction string) {
	//function to change block transaction of the given block ref
	b.Transaction = newTransaction
	b.Cur_Hash = CalculateHash(fmt.Sprintf("%d%s%d%s%s", b.Index, b.Transaction, b.Nonce, b.Prev_Hash, b.Cur_Hash))
}

func VerifyChain(blocks []*Block) bool {
	//function to verify blockchain in case any changes are made. This can be done in two different ways:
	// we are using simple hash verification
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		//checking for current block integrity
		if currentBlock.Cur_Hash != CalculateHash(fmt.Sprintf("%d%s%d%s", currentBlock.Index, currentBlock.Transaction, currentBlock.Nonce, currentBlock.Prev_Hash)) {
			fmt.Printf("chain verification failed at Block %d\n", i)
			return false
		}
		//checking for previous block integrity
		if currentBlock.Prev_Hash != previousBlock.Cur_Hash {
			fmt.Printf("chain verification failed at Block %d\n", i-1)
			return false
		}
	}
	return true
}

func CalculateHash(stringToHash string) string {
	//function for calculating hash of a block
	hashBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashBytes[:])
}
