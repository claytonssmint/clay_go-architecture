package main

import (
	// mysql
	"github.com/claytonssmint/clay_go-architecture/internal/infra/web/server"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	err := server.InitServer()
	if err != nil {
		log.Printf("Error starting server: %s", err)
	}
}
