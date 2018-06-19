package main

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const blocksBucket = "blocks"

// Blockchain holds the blocks that make up the blockchain.
type Blockchain struct {
	db  *badger.DB
	tip []byte
}

// NewBlockchain creates and returns a new Blockchain struct.
func NewBlockchain(DBService *badger.DB) (*Blockchain, error) {
	var tip []byte
	genesis := NewGenesisBlock()

	txn := DBService.NewTransaction(true)
	defer txn.Discard()

	item, err := txn.Get(genesis.Hash)
	if err == badger.ErrKeyNotFound {
		fmt.Printf("Genesis block not found in DB, creating it now.")
		tip = genesis.Hash
		err = txn.Set(genesis.Hash, genesis.Serialize())
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Printf("Genesis block found in DB, setting Blockchain tip to the Genesis hash.")
		tip = item.Key()
	}

	return &Blockchain{
		db:  DBService,
		tip: tip,
	}, nil
}

// AddBlock adds a block to the blockchain.
// func (bc *Blockchain) AddBlock(data string) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	newBlock := NewBlock(data, prevBlock.Hash)
// 	bc.blocks = append(bc.blocks, newBlock)
// }

// NewGenesisBlock creates a new genesis block on the blockchain.
func NewGenesisBlock() *Block {
	return NewBlock("This is the Genesis block", []byte{})
}

// func GetBlockByKey(txn *badger.Txn, key []byte) (*Block, error) {

// }
