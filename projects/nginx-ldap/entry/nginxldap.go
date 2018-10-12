package main

import (
	"log"

	"codelieche.com/authserver"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	authserver.Run()
}
