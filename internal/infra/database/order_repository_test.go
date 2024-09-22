package database

import (
	"database/sql"
	"github.com/claytonssmint/clay_go-architecture/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec(`CREATE TABLE orders (
		id varchar(255) NOT NULL, 
		product VARCHAR(255) NOT NULL, 
		description TEXT NOT NULL, 
		price FLOAT NOT NULL, 
		tax FLOAT NOT NULL, 
		total_price FLOAT NOT NULL, 
		PRIMARY KEY (id)
	)`)
	suite.NoError(err)

	suite.Db = db
}
func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestOrderRepository_SaveSucess() {
	order, err := entity.NewOrder("1", "Product 1", "Description 1", 10.0, 1.0)
	suite.NoError(err)
	suite.NoError(order.CalculateTotalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, product, description, price, tax, total_price FROM orders WHERE id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Product, &orderResult.Description, &orderResult.Price, &orderResult.Tax, &orderResult.TotalPrice)
	suite.NoError(err)

	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Product, orderResult.Product)
	suite.Equal(order.Description, orderResult.Description)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.TotalPrice, orderResult.TotalPrice)
}

func (suite *OrderRepositoryTestSuite) TestOrderRepository_SaveError() {
	order, err := entity.NewOrder("2", "Product 1", "Description 1", 10.0, 1.0)
	suite.NoError(err)
	suite.NoError(order.CalculateTotalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	err = repo.Save(order)
	suite.Error(err)
}
