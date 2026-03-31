package block

import (
	"crypto/sha256"
)

type Block struct {
	HashPrevious []byte
	Data         []byte
	Nonce        int64
	HashCurrent  []byte
}

func ComputeBlock(block *Block) {
	totalData := append(block.HashPrevious, block.Data...)
	totalData = append(totalData, int64ToBytes(block.Nonce)...)
	hash := sha256.Sum256(totalData)
	block.HashCurrent = hash[:]
}

func int64ToBytes(n int64) []byte {
	return []byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	}
}
