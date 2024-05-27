package pow

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"
)

var defaultDifficulty = 16

type ProofOfWork struct {
	mu sync.Mutex

	difficulty int
	block      *Block
}

func New() *ProofOfWork {
	block, err := NewBlock("init", []byte{}).Run(0)
	if err != nil {
		block = &Block{time.Now().Unix(), []byte("init"), []byte{}, []byte{}, 0}
	}
	return &ProofOfWork{difficulty: defaultDifficulty, block: block}
}

func (pow *ProofOfWork) NewTask() (string, int, error) {
	pow.mu.Lock()
	defer pow.mu.Unlock()

	block, err := pow.block.Run(0)
	if err != nil {
		block = &Block{time.Now().Unix(), []byte("init"), []byte{}, []byte{}, 0}
	}
	return fmt.Sprintf("%x", block.Hash), pow.difficulty, nil
}

func (pow *ProofOfWork) Validate(response string) (bool, error) {
	result := Block{}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return false, err
	}
	return result.Validate(pow.difficulty), nil
}

func (block *Block) Validate(difficulty int) bool {
	var hashInt big.Int

	data := block.prepareData(block.Nonce, difficulty)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(setTarget(difficulty)) == -1
}
