package pow

import (
	"crypto/sha256"
	"log"
	"time"
)

// ProofOfWorkClient represents a proof-of-work client part.
type ProofOfWorkClient struct{}

// NewClient return a client's service.
func NewClient() *ProofOfWorkClient {
	return &ProofOfWorkClient{}
}

// Solve returns block with difficulty requested.
func (pow *ProofOfWorkClient) Solve(data string, prevHash []byte, difficulty int) *Block {
	block := Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	newblock, err := block.Run(difficulty)
	if err != nil {
		log.Println("solve error:", err)
		data := block.prepareData(0, difficulty)
		hash := sha256.Sum256(data)
		block.Hash = hash[:]
	}
	return newblock
}
