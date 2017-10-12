package revchain

type Chain struct {
	Blocks []Block
}

// New instantiates a new chain
func (t Chain) New(genesis Block) Chain {
	t.Blocks = make([]Block, 1)
	t.Blocks[0] = genesis

	return t
}

func (t Chain) AddBlock(block Block) Chain {
	t.Blocks = append(t.Blocks, block)

	return t
}

func (t Chain) GetLast() Block {
	return t.Blocks[len(t.Blocks) - 1]
}