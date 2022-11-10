package models

import "time"

type BlockChain struct {
	Blocks_ []*Block
}

func (bc *BlockChain) AddBlock(data BookIssue) {

	prevBlock := bc.Blocks_[len(bc.Blocks_)-1]

	block := CreateBlock(prevBlock, data)

	if ValidBlock(block, prevBlock) {
		bc.Blocks_ = append(bc.Blocks_, block)
	}
}

func CreateBlock(prevBlock *Block, issued BookIssue) *Block {

	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.PreHash = prevBlock.Hash
	block.Data = issued
	block.Timestamp = time.Now().String()
	block.GenerateHash()

	return block
}

func ValidBlock(blo *Block, preblo *Block) bool {

	if preblo.Hash != blo.PreHash {
		return false
	}

	if !blo.ValidateHash(blo.Hash) {
		return false
	}

	if preblo.Pos + 1 != blo.Pos {
		return false
	}


 	return true
}


