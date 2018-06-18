package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

const blocksBucket = "blocks"

// Blockchain holds the blocks that make up the blockchain.
type Blockchain struct {
	db *badger.DB
}

// NewBlockchain creates and returns a new Blockchain struct.
func NewBlockchain(DBService *badger.DB) *Blockchain {
	//var tip []byte
	genesis := NewGenesisBlock()

	fmt.Printf("hash is %+v \n \n", genesis.Hash)
	fmt.Printf("serialized block is %+v \n \n", genesis.Serialize())

	txn := DBService.NewTransaction(true)
	defer txn.Discard()

	item, err := txn.Get(genesis.Hash)
	if err != nil {
		log.Fatal("Urgh can't find it")
	}

	fmt.Printf("item %+v", item)

	return &Blockchain{
		db: DBService,
	}
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
