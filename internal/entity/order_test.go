package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnAnErrorWhenEnteringAnEmptyID(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid order id")
}

func Test_ShouldReturnAnErrorWhenEnteringAnEmptyProduct(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.IsValid(), "invalid order product")
}

func Test_ShouldReturnAnErrorWhenEnteringAnEmptyDescription(t *testing.T) {
	order := Order{ID: "1", Product: "Product"}
	assert.Error(t, order.IsValid(), "invalid order description")
}

func Test_ShouldReturnAnErrorWhenEnteringAnInvalidPrice(t *testing.T) {
	order := Order{ID: "1", Product: "Product", Description: "Description"}
	assert.Error(t, order.IsValid(), "invalid order price")
}

func Test_ShouldReturnAnErrorWhenEnteringAnInvalidTax(t *testing.T) {
	order := Order{ID: "1", Product: "Product", Description: "Description", Price: 4.000}
	assert.Error(t, order.IsValid(), "invalid order tax")
}

func Test_ShouldReturnNilWhenEnteringAValidOrder(t *testing.T) {
	order := Order{ID: "1", Product: "Laptop", Description: "High-performance laptop 16GB RAM", Price: 4.000, Tax: 1}
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "Laptop", order.Product)
	assert.Equal(t, "High-performance laptop 16GB RAM", order.Description)
	assert.Equal(t, 4.000, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	assert.Nil(t, order.IsValid())
}
