package main

import (
	"log"

	"github.com/adarsh-jaiss/broker/broker/client"
)

func main() {
	s := client.NewServer(&client.Server{
		PORT: "8001",
		HOST: "localhost",
	})

	log.Println("starting a server at:", s.PORT)
	s.InitTCPServer()
}
