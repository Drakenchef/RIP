package main

import (
	"github.com/drakenchef/RIP/internal/api"
	"log"
)

func main() {
	log.Println("Server started")
	api.StartServer()
	log.Println("Server shutdown")
}

//selfcommit
