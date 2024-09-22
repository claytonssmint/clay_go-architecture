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

func Test_GivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("1", "Laptop", "High-performance laptop 16GB RAM", 4.000, 1)
	assert.Nil(t, err)
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "Laptop", order.Product)
	assert.Equal(t, "High-performance laptop 16GB RAM", order.Description)
	assert.Equal(t, 4.000, order.Price)
	assert.Equal(t, 1.0, order.Tax)
}

func Test_GivenAnInvalidParams_WhenICallNewOrderFunc_ThenIShouldReceiveAnError(t *testing.T) {
	order, err := NewOrder("1", "Laptop", "High-performance laptop 16GB RAM", 0, 0)
	assert.NotNil(t, err)
	assert.Nil(t, order)
}

func Test_GivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("1", "Laptop", "High-performance laptop 16GB RAM", 4.000, 1)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateTotalPrice())
	assert.Equal(t, 5.000, order.TotalPrice)
}
