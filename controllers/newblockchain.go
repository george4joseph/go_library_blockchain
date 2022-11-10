package controllers

import "github.com/george4joseph/go_library_blockchain/models"

func NewBlockChain() *models.BlockChain {

	return &models.BlockChain{Blocks_: []*models.Block{GenesisBlock()}}
}

func GenesisBlock() *models.Block {
	return models.CreateBlock(&models.Block{}, models.BookIssue{IsGenesis: true})
}
