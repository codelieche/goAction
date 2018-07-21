package main

import (
	"goAction/tutorial/simplechain"
	"log"
	"fmt"
	"time"
)

func main () {
	blockChain := simplechain.NewBlockChain()
	log.Println(blockChain)
	for i := 1; i < 10; i++ {
		data := fmt.Sprintf("Block %v", i)
		blockChain.SendData(data)
		log.Println(blockChain)
	}

	time.Sleep(time.Duration(time.Second))
	blockChain.Print()
}