package main

// Blockchain holds the blocks that make up the blockchain.
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain creates and returns a new Blockchain struct.
func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{
			NewGenesisBlock(),
		},
	}
}

// AddBlock adds a block to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock creates a new genesis block on the blockchain.
func NewGenesisBlock() *Block {
	return NewBlock("This is the Genesis block", []byte{})
}
