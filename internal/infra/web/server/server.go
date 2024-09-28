package server

import (
	"database/sql"
	"fmt"
	"github.com/claytonssmint/clay_go-architecture/configs"
	"github.com/claytonssmint/clay_go-architecture/internal/infra/database"
	"github.com/claytonssmint/clay_go-architecture/internal/infra/web"
	"github.com/claytonssmint/clay_go-architecture/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func InitServer() error {
	configs, err := configs.LoadConfig("cmd/application")
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações: %v", err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCreateOrderUseCase(orderRepository)
	webOrderHandler := web.NewWebOrderHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	setupRoutes(r, webOrderHandler)

	log.Printf("Starting server on %s\n", configs.WebServerPort)
	return http.ListenAndServe(configs.WebServerPort, r)
}

func setupRoutes(r chi.Router, webOrderHandler *web.WebOrderHandler) {

	r.Post("/order", webOrderHandler.PostOrder)
	r.Get("/order/{id}", webOrderHandler.GetOrderID)
}
