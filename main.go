package main

import (
	"blockchain/block"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var hashStrength uint8
	var blocksAmount int8
	fmt.Println("Enter number of blocks to put in a blockchain (must be enough data): ")
	if _, err := fmt.Scan(&blocksAmount); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Enter hash strength (number of 0's as a prefix) (not larger than 12): ")
	if _, err := fmt.Scan(&hashStrength); err != nil || hashStrength > 12 {
		fmt.Println("You've entered something wrong...")
		return
	}
	prefix := strings.Repeat("0", int(hashStrength))

	fmt.Println("Scanned successfully!")
	data, _ := os.ReadFile("block.1")
	blockChain := make([]block.Block, blocksAmount)
	genesis := block.Block{
		[]byte{},
		data,
		0,
		[]byte{},
	}
	block.ComputeBlock(&genesis)
	findNonce(&genesis, prefix)
	blockChain[0] = genesis
	var newBlock block.Block
	for i := 1; i < int(blocksAmount); i++ {
		data, err := os.ReadFile("block." + strconv.Itoa(i+1))
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Looks like there isn't enough data for the blocks...")
			return
		}
		newBlock.Data = data
		newBlock.Nonce = 0
		newBlock.HashPrevious = blockChain[i-1].HashCurrent
		findNonce(&newBlock, prefix)
		blockChain[i] = newBlock
	}

	for i, bl := range blockChain {
		fmt.Println("Block number", i+1)
		fmt.Println("\tNonce:", bl.Nonce)
		fmt.Println("\tBlock's Hash:", hex.EncodeToString(bl.HashCurrent))
	}
}
func findNonce(bl *block.Block, prefix string) {
	for !strings.HasPrefix(hex.EncodeToString(bl.HashCurrent), prefix) {
		block.ComputeBlock(bl)
		bl.Nonce++
	}
}
