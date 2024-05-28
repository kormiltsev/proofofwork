package pow

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/kormiltsev/proofofwork/config"
)

// ProofOfWork implements proof of work scenario for one block.
type ProofOfWork struct {
	mu  sync.Mutex
	log log15.Logger

	difficulty int
	block      *Block

	db    map[string]bool
	olddb map[string]bool
	limit int
}

// New returns service of PoW.
func New() *ProofOfWork {
	block, err := NewBlock("init", []byte{}).Run(0)
	if err != nil {
		block = &Block{time.Now().Unix(), []byte("init"), []byte{}, []byte{}, 0}
	}
	return &ProofOfWork{log: log15.New("controller", "job"), difficulty: config.Difficulty, block: block, db: map[string]bool{}, olddb: map[string]bool{}, limit: config.Limit}

}

// NewTask return hash and difficulty requested.
func (pow *ProofOfWork) NewTask() (string, int, error) {

	block, err := pow.block.Run(0)
	if err != nil {
		block = &Block{time.Now().Unix(), []byte("init"), []byte{}, []byte{}, 0}
	}
	pow.addKey(block.Hash)

	return fmt.Sprintf("%x", block.Hash), pow.difficulty, nil
}

// Validate check the answer. Error in case of unmarshal error only.
func (pow *ProofOfWork) Validate(response string) (bool, error) {
	result := Block{}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return false, err
	}

	if !result.Validate(pow.difficulty) {
		return false, nil
	}
	return pow.checkKey(result.PrevBlockHash), nil
}

// Validate for block implements validation for the block provided.
func (block *Block) Validate(difficulty int) bool {
	var hashInt big.Int

	data := block.prepareData(block.Nonce, difficulty)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(setTarget(difficulty)) == -1
}

// save hash to cache
func (pow *ProofOfWork) addKey(hash []byte) {
	pow.mu.Lock()
	defer pow.mu.Unlock()
	pow.db[fmt.Sprintf("%x", hash)] = true
}

// check if task was created
func (pow *ProofOfWork) checkKey(hash []byte) bool {
	pow.mu.Lock()
	defer pow.mu.Unlock()

	if len(pow.db) >= pow.limit {
		pow.olddb = pow.db
		pow.db = map[string]bool{}
	}

	key := string(hash)

	if pow.db[key] {
		delete(pow.db, key)
		return true
	}

	if pow.olddb[key] {
		delete(pow.olddb, key)
		return true
	}
	return false
}
