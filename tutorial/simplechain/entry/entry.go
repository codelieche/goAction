package main

import (
	"goAction/tutorial/simplechain"
	"log"
	"fmt"
)

func main () {
	root := simplechain.GenerateGenesisBlock()
	log.Println(root)

	prevBlock := root
	for i := 1; i < 10; i++ {
		data := fmt.Sprintf("Block %v", i)
		nextBlock := prevBlock.GeneraterNextBlock(data)
		prevBlock =  nextBlock

		log.Println(nextBlock)
	}
}