package web

import (
	"encoding/json"
	"github.com/claytonssmint/clay_go-architecture/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type WebOrderHandler struct {
	usecase usecase.UseCasePort
}

func NewWebOrderHandler(uc usecase.UseCasePort) *WebOrderHandler {
	return &WebOrderHandler{usecase: uc}
}

func (handler *WebOrderHandler) PostOrder(w http.ResponseWriter, r *http.Request) {
	var dtoResp usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dtoResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dtoResp.ID = uuid.New().String()

	output, err := handler.usecase.CreateUseCase(dtoResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *WebOrderHandler) GetOrderID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, err := handler.usecase.FindByIDUseCase(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
