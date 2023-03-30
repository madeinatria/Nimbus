package main

import (
	"log"

	"github.com/madeinatria/Nimbus/cmd/datastore"
	"github.com/madeinatria/Nimbus/cmd/server"
)

func main() {
	log.Println(datastore.Db)
	server.Start()
}
