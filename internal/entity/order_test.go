package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnAnErrorWhenEnteringAnEmptyID(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid order id")
}
