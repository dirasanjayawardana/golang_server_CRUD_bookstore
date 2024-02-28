package repositories

import (
	"golang_server_bookstore/internals/models"

	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	*sqlx.DB
}

func InitBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db}
}

func (item BookRepo) FindAll() ([]models.BookModel, error) {
	query := "SELECT * FROM books"
	result := []models.BookModel{}
	// .select akan mereturnkan hasil dari query
	if err := item.Select(&result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (item *BookRepo) SaveBook(body models.BookModel) error {
	query := "INSERT INTO books(title, description, author) VALUES (?, ?, ?)"
	// .exce --> tidak mereturn rows (data dari database)
	if _, err := item.Exec(query, body.Title, body.Description, body.Author); err != nil {
		return err
	}
	return nil
}
