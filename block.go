package main

import (
	"bytes"
	"encoding/gob"
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

// DeserializeBlock deserializes a serialized block and returns it.
func DeserializeBlock(serializedBlock []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(serializedBlock))
	decoder.Decode(&block)

	return &block
}

// NewBlock creates and returns a new Block struct.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Serialize serializes a block into an array of bytes.
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	encoder.Encode(b)

	return result.Bytes()
}
