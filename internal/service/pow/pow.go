package pow

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/kormiltsev/proofofwork/internal/utils"
)

// Block is for one time validation.
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock prepare and return a new block.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	block, err := block.Run(0)
	if err != nil {
		data := block.prepareData(0, 0)
		hash := sha256.Sum256(data)
		block.Hash = hash[:]
	}
	return block
}

// Run performs a proof-of-work.
func (block *Block) Run(difficulty int) (*Block, error) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := block.prepareData(nonce, difficulty)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(setTarget(difficulty)) == -1 {
			fmt.Printf("\n🟢 %x\n", hash)
			block.Hash = hash[:]
			block.Nonce = nonce
			return block, nil
		} else {
			nonce++
		}
	}

	return nil, fmt.Errorf("not solved")
}

func setTarget(targetBits int) *big.Int {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return target
}

func (block *Block) prepareData(nonce, difficulty int) []byte {
	data := bytes.Join(
		[][]byte{
			block.PrevBlockHash,
			block.Data,
			utils.IntToHex(block.Timestamp),
			utils.IntToHex(setTarget(difficulty).Int64()),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}
