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
