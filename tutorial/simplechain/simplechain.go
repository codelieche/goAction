package simplechain

func NewBlockChain() *BlockChain {
	// 创建一个区块链
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.Append(genesisBlock)
	return &blockChain
}
