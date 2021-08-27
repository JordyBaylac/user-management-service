package main

import (
	"log"

	"github.com/JordyBaylac/user-management-service/api"
)

func main() {
	log.Fatal(api.Start())
}
