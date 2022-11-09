package models


type BlockChain struct {

	blocks []*Block
}

func (bc *BlockChain)AddBlock(data BookIssue) {

	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := createBlock(prevBlock, data)


}