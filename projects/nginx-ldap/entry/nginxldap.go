package main

import (
	"log"

	"goAction/projects/nginx-ldap/web/authserver"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	authserver.Run()
}
