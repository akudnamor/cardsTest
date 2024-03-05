package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS products(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON products(name);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetProductWithOffset(limit, offset int) ([]Product, error) {
	const op = "storage.sqlite.GetProductWithOffset"
	stmt, err := s.db.Prepare("SELECT * FROM products LIMIT (?) OFFSET (?)")
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) GetAllProduct() ([]Product, error) {
	const op = "storage.sqlite.GetAllProduct"
	stmt, err := s.db.Prepare("SELECT id, name, price FROM products")
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) AddProduct(id int, name string, price float64) error {
	const op = "storage.sqlite.AddProduct"
	stmt, err := s.db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	_, err = stmt.Exec(id, name, price)
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
