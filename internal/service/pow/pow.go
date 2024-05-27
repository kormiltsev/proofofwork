package pow

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/kormiltsev/proofofwork/internal/utils"
)

// ProofOfWork represents a proof-of-work service.
type ProofOfWorkClient struct{}

// Block is for one time validation.
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

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

func NewClient() *ProofOfWorkClient {
	return &ProofOfWorkClient{}
}

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

// Run performs a proof-of-work
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
