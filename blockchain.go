package main

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	block []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock (data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
