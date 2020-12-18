package main

import (
	"fmt"
	"log"
	"time"

	"github.com/codelieche/goAction/tutorial/demo/simplechain"
)

func main() {
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
