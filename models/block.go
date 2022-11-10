package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	Pos       int
	Data      BookIssue
	Timestamp string
	Hash      string
	PreHash   string
}

func (b *Block) GenerateHash() {
	bytes, _ := json.Marshal(b.Data)

	data := string(b.Pos) + b.Timestamp + string(bytes) + b.PreHash

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))

}


func (b *Block) ValidateHash(hash string) bool {
	b.GenerateHash()
	if b.Hash != hash {
		return false
	}
	return true
}
