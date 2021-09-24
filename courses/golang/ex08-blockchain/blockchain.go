package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          string
	PrevBlockHash string
	Hash          string
}

type Blockchain struct {
	blocks []*Block
}

func (block Block) String() string {
	return fmt.Sprintf("Prev. hash: %s\nData: %s\nHash: %s\n", block.PrevBlockHash, block.Data, block.Hash)
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{[]byte(b.PrevBlockHash), []byte(b.Data), timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{time.Now().UnixNano(), data, prevBlockHash, ""}
	block.setHash()

	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
