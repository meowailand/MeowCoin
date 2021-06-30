package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

// the singleton pattern is a software design pattern that restricts the instantiation of a class to one "single" instance
var b *blockchain

func GetBlockChain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}
