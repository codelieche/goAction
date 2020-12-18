package simplechain

//生成创世区块

func GenerateGenesisBlock() *Block{
	// 生成创世区块
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	rootBlock := preBlock.GeneraterNextBlock("Genesis Block：创世区块")
	return rootBlock
}

