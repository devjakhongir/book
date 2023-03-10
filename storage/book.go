package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func InsertBook(db *sql.DB, book models.CreateBook) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO books (
			id,
			name,
			price,
			description,
			updated_at
		) VALUES ($1, $2, $3, $4, now())
	`

	_, err := db.Exec(query,
		id,
		book.Name,
		book.Price,
		book.Description,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdBook(db *sql.DB, req models.BookPrimeryKey) (models.Book, error) {

	var (
		book models.Book
	)

	query := `
		SELECT
			id,
			name,
			price,
			description,
			created_at,
			updated_at
		FROM books WHERE id = $1
	`

	err := db.QueryRow(query, req.Id).Scan(
		&book.Id,
		&book.Name,
		&book.Price,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func GetListBook(db *sql.DB, req models.GetListBookRequest) (models.GetListBookResponse, error) {

	var (
		resp   models.GetListBookResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			price,
			description,
			created_at,
			updated_at
		FROM books
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := db.Query(query)
	if err != nil {
		return models.GetListBookResponse{}, err
	}

	for rows.Next() {
		var book models.Book

		err = rows.Scan(
			&resp.Count,
			&book.Id,
			&book.Name,
			&book.Price,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return models.GetListBookResponse{}, err
		}

		resp.Books = append(resp.Books, book)
	}

	return resp, nil
}

func UpdateBook(db *sql.DB, book models.Book) (int64, error) {

	query := `
		UPDATE 
			books 
		SET 
			name = $2,
			price = $3,
			description = $4,
			updated_at = now()
		WHERE id = $1
	`

	result, err := db.Exec(query,
		book.Id,
		book.Name,
		book.Price,
		book.Description,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteBook(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
