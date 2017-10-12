package revchain

import (
	"fmt"
	"crypto/sha256"
	"strconv"
	"encoding/json"
	"time"
)

type Block struct {
	PrevHash string
	Index int
	Kekspace interface{}
	Timestamp int64
	Data []byte
	Hash []byte
}

func (b Block) New(ks interface{}, data []byte, pHash string, index int) Block {
	b.PrevHash = pHash
	b.Index = index
	b.Timestamp = time.Now().Unix()
	b.Kekspace = ks
	b.Data = data
	b.Hash = b.GenHash(b)

	return b
}

func (block Block) isHashValid(hash string) bool {
	validStart := hash[:2]
	return validStart == "00"
}

func (block Block) GenHash(b Block) []byte {
	sha := sha256.New()
	index := strconv.Itoa(b.Index)
	auth, _ := json.Marshal(b.Kekspace)
	newAuth := fmt.Sprint(auth)
	nonce := 0
	hash := "aaaa"
	data := fmt.Sprintf("%x", block.Data)

	for !b.isHashValid(hash) {
		nonce++
		sha.Write([]byte(index + newAuth + b.PrevHash + data + strconv.Itoa(nonce)))
		newHash := sha.Sum(nil)
		hash = fmt.Sprintf("%x", newHash)
	}

	return sha.Sum(nil)
}

func (b Block) ValidateHash() bool {
	return true
}

func (b Block) HashString() string {
	return fmt.Sprintf("%x",b.GenHash(b))
}

