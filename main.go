package main

import (
	bc "blockchain/block"
	"fmt"
)

func main() {
	var hashStrength uint
	var blocksAmount uint8
	fmt.Println("Enter number of blocks to put in a blockchain (must be enough data): ")
	if _, err := fmt.Scan(&blocksAmount); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Enter hash strength (number of 0 bits as a prefix) (do not pick too large): ")
	if _, err := fmt.Scan(&hashStrength); err != nil || hashStrength > 128 {
		fmt.Println("You've entered something wrong...")
		return
	}
	fmt.Println("Scanned successfully!")

	bc.View(blocksAmount, hashStrength)
}
