package entity

import "errors"

type Order struct {
	ID          string
	Product     string
	Description string
	Price       float64
	Tax         float64
	TotalPrice  float64
}

func NewOrder(id, product, description string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:          id,
		Product:     product,
		Description: description,
		Price:       price,
		Tax:         tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid order id")
	}
	if o.Product == "" {
		return errors.New("invalid order product")
	}
	if o.Price <= 0 {
		return errors.New("invalid order price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid order tax")
	}
	return nil
}

func (o *Order) CalculateTotalPrice() error {
	o.TotalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
