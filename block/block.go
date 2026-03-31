package block

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

type Block struct {
	HashPrevious []byte `json:"hash_previous"`
	Data         []byte `json:"data"`
	Nonce        int64  `json:"nonce"`
	HashCurrent  []byte `json:"hash_current"`
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

func ValidateBlock(block *Block) bool {
	totalData := append(block.HashPrevious, block.Data...)
	totalData = append(totalData, int64ToBytes(block.Nonce)...)
	hash := sha256.Sum256(totalData)
	return bytes.Equal(hash[:], block.HashCurrent)
}

func FindNonce(bl *Block, hashStrength uint) {
	for !ValidHash(bl.HashCurrent, hashStrength) {
		bl.Nonce++
		ComputeBlock(bl)
	}
}

func ValidHash(hash []byte, difficulty uint) bool {
	target := new(big.Int).Lsh(big.NewInt(1), 256-difficulty)
	hashInt := new(big.Int).SetBytes(hash)
	return hashInt.Cmp(target) == -1
}
