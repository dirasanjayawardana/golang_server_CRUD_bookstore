package models

// untuk arahan konversi dari body berupa json atau form ke database atau sebaliknya
type AuthModel struct {
	Id       int    `db:"id"`
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password" form:"password"`
}
