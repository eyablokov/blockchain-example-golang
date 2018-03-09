package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/*
Block keeps block headers.
	Timestamp is the current timestamp (when the block is created).
	Data is the actual valuable information containing in the block.
	PrevBlockHash is the hash of the previous block.
	Hash is the hash of the block.
*/
type Block struct {
	Timestamp     int64
	Data          []bytes
	PrevBlockHash []bytes
	Hash          []bytes
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
