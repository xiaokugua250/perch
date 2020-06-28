package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"sync"
	"time"
)

/**
Index is the position of the data record in the blockchain
Timestamp is automatically determined and is the time the data is written
BPM or beats per minute, is your pulse rate
Hash is a SHA256 identifier representing this data record
PrevHash is the SHA256 identifier of the previous record in the chain
*/
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PreHash   string
}

var BlockChain []Block
var BlockChainServer chan []Block //// bcServer handles incoming concurrent Blocks
var mutex = &sync.Mutex{}

/**
初始块创建
*/
func GenInitialChainBlock() error {
	genesisBlock := Block{}
	genesisBlock = Block{0, time.Now().String(), 0, cacalueHash(genesisBlock), ""} //初始块
	spew.Dump(genesisBlock)
	mutex.Lock()
	BlockChain = append(BlockChain, genesisBlock)
	mutex.Unlock()
	BlockChain = append(BlockChain)
	return nil
}

/**
计算hash 块
*/
func cacalueHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

/**
根据旧块生产新的hash块
*/
func GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = cacalueHash(newBlock)
	return newBlock, nil

}

/**
检查块是否合法有效
*/
func IsBlockValied(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PreHash {
		return false
	}
	if cacalueHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(BlockChain) {
		BlockChain = newBlocks
	}
}
