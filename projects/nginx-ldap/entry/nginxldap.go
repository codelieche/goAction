package main

import (
	"log"

	"codelieche.com/ldaplib"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ldaplib.Auth("xxx", "xxx")
}
