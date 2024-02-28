// repository --> berupa struct, berisi method untuk query ke database

package repositories

import (
	"golang_server_bookstore/internals/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	*sqlx.DB
}

func InitAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (item *AuthRepo) FindByEmail(body models.AuthModel) ([]models.AuthModel, error) {
	query := "SELECT * FROM users WHERE email = ?"
	result := []models.AuthModel{}
	// .select akan mereturnkan hasil
	if err := item.Select(&result, query, body.Email); err != nil {
		return nil, err
	}
	// untuk custom query --> item.Query()
	return result, nil
}

func (item *AuthRepo) SaveUser(body models.AuthModel) error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	if _, err := item.Exec(query, body.Email, body.Password); err != nil {
		return err
	}
	return nil
}
