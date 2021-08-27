package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/JordyBaylac/user-management-service/api"
)

const defaultPort = 8080

func main() {
	app := api.Setup()
	address := getServerAddress()

	log.Fatal(app.Listen(address))
}

// getServerAddress reads PORT from environment or default
func getServerAddress() string {
	port := defaultPort

	portEnv := os.Getenv("PORT")
	if number, err := strconv.Atoi(portEnv); err == nil {
		port = number
	}

	return fmt.Sprintf("0.0.0.0:%d", port)
}
