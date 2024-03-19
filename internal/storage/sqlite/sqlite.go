package sqlite

import (
	"BOARD/internal/storage"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

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

// it can be not working
func (s *Storage) GetProductWithOffset(limit, offset int) ([]storage.Product, error) {
	const op = "storage.sqlite.GetProductWithOffset"
	stmt, err := s.db.Prepare("SELECT id, name, price FROM products LIMIT (?) OFFSET (?)")
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}

	var products []storage.Product
	for rows.Next() {
		var product storage.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) GetAllProduct() ([]storage.Product, error) {
	const op = "storage.sqlite.GetAllProduct"
	stmt, err := s.db.Prepare("SELECT id, name, price FROM products")
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}

	var products []storage.Product
	for rows.Next() {
		var product storage.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) GetProductByID(id int) (storage.Product, error) {
	const op = "storage.sqlite.GetAllProduct"
	stmt, err := s.db.Prepare("SELECT id, name, price FROM products where id = (?)")
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRow(id)

	var product storage.Product
	err = row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Errorf("%s: %w", op, err)
	}
	return product, nil
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
