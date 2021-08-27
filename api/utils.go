package api

import (
	"fmt"
	"os"
	"strconv"
)

// getServerAddress reads PORT from environment or default to 8080
func getServerAddress() string {
	port := defaultPort

	portEnv := os.Getenv("PORT")
	if number, err := strconv.Atoi(portEnv); err == nil {
		port = number
	}

	return fmt.Sprintf("0.0.0.0:%d", port)
}
