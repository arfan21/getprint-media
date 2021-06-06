package main

import (
	"log"

	"github.com/arfan21/getprint-media/server"
	_ "github.com/joho/godotenv"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
