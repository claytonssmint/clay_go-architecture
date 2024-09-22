package usecase

import (
	"github.com/claytonssmint/clay_go-architecture/internal/entity"
)

type UseCasePort interface {
	Apply(input OrderInputDTO) (OrderOutputDTO, error)
}

type OrderInputDTO struct {
	ID          string  `json:"id"`
	Product     string  `json:"product"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Tax         float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID          string  `json:"id"`
	Product     string  `json:"product"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Tax         float64 `json:"tax"`
	TotalPrice  float64 `json:"total_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepositoryInterface) UseCasePort {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *CreateOrderUseCase) Apply(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:          input.ID,
		Product:     input.Product,
		Description: input.Description,
		Price:       input.Price,
		Tax:         input.Tax,
	}
	order.CalculateTotalPrice()
	if err := u.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}
	outputDTO := OrderOutputDTO{
		ID:          order.ID,
		Product:     order.Product,
		Description: order.Description,
		Price:       order.Price,
		Tax:         order.Tax,
		TotalPrice:  order.TotalPrice,
	}
	return outputDTO, nil
}
