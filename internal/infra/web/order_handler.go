package web

import (
	"encoding/json"
	"github.com/claytonssmint/clay_go-architecture/internal/usecase"
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

	output, err := handler.usecase.Apply(dtoResp)
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
