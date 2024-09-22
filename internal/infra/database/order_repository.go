package database

import (
	"database/sql"
	"github.com/claytonssmint/clay_go-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, product, description, price, tax, total_price) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Product, order.Description, order.Price, order.Tax, order.TotalPrice)
	if err != nil {
		return err
	}
	return nil
}
