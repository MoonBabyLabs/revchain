package revchain

import (
	"github.com/MoonBabyLabs/kekstore"
	"github.com/MoonBabyLabs/kekspace"
)

type Chain struct {
	store kekstore.Storer
	Blocks []Block `json:"blocks"`
	CurHash string `json:"current_hash"`
	Index int `json:"index"`
	Structure map[string]string `json:"structure"`
}

type ChainMaker interface {
	New(itemId string, data interface{}) (ChainMaker, error)
	GetHashString() string
	AddBlock(itemId string, data interface{}) (ChainMaker, error)
	Load(path string) (ChainMaker, error)
	Delete(id string) error
	GetBlocks() []Block
}

const KEK_PATH = "k/"

func (t Chain) Delete(id string) error {
	return t.store.Delete(KEK_PATH + id + ".kek")
}

func (t Chain) GetBlocks() []Block {
	return t.Blocks
}

func (t Chain) AddBlock(id string, data interface{}) (ChainMaker, error) {
	space, spErr := kekspace.Kekspace{}.Load()

	if spErr != nil {
		return t, spErr
	}

	block := Block{}.New(space.KekId, data, t.CurHash, t.Index + 1)
	t.CurHash = block.Hash
	t.Index = block.Index
	t.Blocks = append(t.Blocks, block)
	t.store.Save(KEK_PATH + id + ".kek", t)

	return t, nil
}

func (t Chain) SetStore(store kekstore.Storer) Chain {
	t.store = store

	return t
}

func (t Chain) New(id string, data interface{}) (ChainMaker, error) {
	space, spaceErr := kekspace.Kekspace{}.Load()

	if spaceErr != nil {
		return t, spaceErr
	}

	if t.store == nil {
		t.store = kekstore.Store{}
	}

	block := Block{}.New(space.KekId, data, "", 0)
	t.Blocks = []Block{block}
	t.Index = 0
	t.CurHash = block.Hash
	t.store.Save(KEK_PATH + id + ".kek", t)

	return t, nil
}

func (t Chain) GetHashString() string {
	return t.CurHash
}

func (t Chain) Load(id string) (ChainMaker, error) {
	if t.store == nil {
		t.store = kekstore.Store{}
	}

	loadErr := t.store.Load(KEK_PATH + id + ".kek", &t)

	return t, loadErr
}