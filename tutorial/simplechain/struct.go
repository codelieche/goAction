package simplechain

import (
	"time"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Index int64  // 区块编号：自增
	TimeStamp int64 // 区块生成时间
	PrevHash string // 上一个区块的哈希值
	Hash string // 当前区块的哈希值
	Data string // 区块的数据
}

func (b *Block) CalculateHash() string {
	// 计算区块的哈希值
	blockData := fmt.Sprintf("%d%d%s%s", b.Index, b.TimeStamp, b.PrevHash, b.Data);
	// 使用sha256计算Hash值
	hashBytes := sha256.Sum256([]byte(blockData))
	// 返回字符串16进制
	return hex.EncodeToString(hashBytes[:]);
}

func (b *Block) GeneraterNextBlock(data string) *Block {
	nextBlock := Block{}
	nextBlock.Index = b.Index + 1
	nextBlock.PrevHash = b.Hash
	nextBlock.TimeStamp = time.Now().UTC().Unix()
	nextBlock.Data = data
	// 计算哈希值
	nextBlock.Hash = nextBlock.CalculateHash()
	return &nextBlock
}

// 区块链
type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) Append(b *Block) error {
	// 往区块链中添加区块
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, b)
		return nil
	}
	// 添加前需要验证
	if bc.isValid(b) {
		bc.Blocks = append(bc.Blocks, b)
		return nil
	}else{
		return fmt.Errorf("区块校验不通过")
	}
}

func (bc *BlockChain) isValid(new *Block) bool {
	old := *bc.Blocks[len(bc.Blocks) - 1]
	if old.Index + 1 != new.Index {
		return false
	}
	if old.Hash != old.PrevHash {
		return false
	}
	if new.CalculateHash() != new.Hash {
		return false
	}
	return true
}

func (bc *BlockChain) SendData(data string){

	// 直接通过字符串，加入区块
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := prevBlock.GeneraterNextBlock(data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *BlockChain) Print(){
	for _, block := range bc.Blocks {
		fmt.Printf("区块: %d\n", block.Index)
		fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Println()
	}
}



