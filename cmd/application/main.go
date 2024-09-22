package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/claytonssmint/clay_go-architecture/configs"
	"github.com/claytonssmint/clay_go-architecture/internal/infra/database"
	"github.com/claytonssmint/clay_go-architecture/internal/infra/web"
	"github.com/claytonssmint/clay_go-architecture/internal/infra/web/server"
	"github.com/claytonssmint/clay_go-architecture/internal/usecase"
	"log"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig("cmd/application")
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações: %v", err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCreateOrderUseCase(orderRepository)

	webserver := server.NewWebServer(configs.WebServerPort)
	webOrderHandler := web.NewWebOrderHandler(uc)

	webserver.AddHandler("/order", webOrderHandler.PostOrder)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start()
}
