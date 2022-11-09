package models


type BlockChain struct {

	blocks []*Block
}

func (bc *BlockChain) AddBlock(data BookIssue) {

	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

