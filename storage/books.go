package storage

import (
	"database/sql"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

type Book struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	AuthorName string    `json:"author_name"`
	Price      float64   `json:"price"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

func (b *DBManager) Create(book *Book) (*Book, error) {
	query := `
		INSERT INTO books(
			title,
			author_name,
			price,
			amount
		) VALUES ($1,$2,$3,$4)
		RETURNING id, title, author_name, price, amount, created_at`

	row := b.db.QueryRow(query,
		book.Title,
		book.AuthorName,
		book.Price,
		book.Amount,
	)
	var result Book
	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.AuthorName,
		&result.Price,
		&result.Amount,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err 
	}
	return &result, nil 
}

func (b *DBManager) Get(id int64) (*Book, error) {
	var result Book

	query := `
		SELECT
			id,
			title,
			author_name,
			price,
			amount,
			created_at
		FROM books
		WHERE id = $1`
	row := b.db.QueryRow(query,id) 
	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.AuthorName,
		&result.Price,
		&result.Amount,
		&result.CreatedAt,
	)
	if err != nil {
		return nil ,err 
	}
	return &result, nil 

}

func (b *DBManager) Update(book *Book) (*Book, error) {
	query := `
		UPDATE books SET
			title = $1,
			author_name = $2,
			price = $3,
			amount = $4
		WHERE id = $5
		RETURNING id, title, author_name, price, amount, created_at
		`
	row := b.db.QueryRow(query,
		book.Title,
		book.AuthorName,
		book.Price,
		book.Amount,
		book.Id,
	)

	var result Book 
	
	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.AuthorName,
		&result.Price,
		&result.Amount,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err 
	}

	return &result, nil 

}

func (b *DBManager) Delete(id int64) error {
	query := `DELETE FROM books WHERE id = $1`

	result, err := b.db.Exec(query,id)
	if err != nil {
		return err 
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err 
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil 
}