package revchain

import (
	"fmt"
	"crypto/sha256"
	"strconv"
	"encoding/json"
	"time"
)

type BlockData struct {
	Add map[string]interface{} `json:"add"`
	Modify map[string]interface{} `json:"modify"`
	Delete map[string]interface{} `json:"delete"`
}

type Signature struct {
	Token string `json:"token"`
	Source string `json:"string"`
}

type Block struct {
	PrevHash string `json:"prev_hash"`
	Index int `json:"index"`
	Kekspace interface{} `json:"space"`
	Timestamp int64 `json:"timestamp"`
	Data BlockData `json:"data"`
	Hash string `json:"hash"`
	Signature Signature `json:"signature"`
}

func (b Block) New(ks interface{}, addData map[string]interface{}, modifyData map[string]interface{}, deleteData map[string]interface{}, pHash string, index int) Block {
	b.PrevHash = pHash
	b.Index = index + 1
	b.Timestamp = time.Now().Unix()
	b.Kekspace = ks
	b.Data = BlockData{Add:addData, Modify:modifyData, Delete:deleteData}
	b.Hash = b.HashString()

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

