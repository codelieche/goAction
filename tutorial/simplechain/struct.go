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

