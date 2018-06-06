package main

import (
	"time"
)

type (
	// Block holds transaction data and is added a single Block to the blockchain.
	Block struct {
		Timestamp     int64
		Data          []byte
		PrevBlockHash []byte
		Nonce         int
		Hash          []byte
	}
)

// NewBlock creates and returns a new Block struct.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}
