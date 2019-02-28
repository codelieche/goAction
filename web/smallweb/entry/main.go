package main

import (
	"goAction/web/smallweb"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	smallweb.RunServer()
}
