package repository

import (
	backend "cmd/main.go"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user backend.User) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash, phone) VALUES ($1,$2,$3,$4) RETURNING id",userTable )
	row := r.db.QueryRow(query,user.Name, user.Username, user.Password, user.Phone )
	if err:= row.Scan(&id);err !=nil{
		return 0, err
	}	
	return id ,nil
}
func (r *AuthPostgres)  GetUser(username,password string) (backend.User, error) {
	var user backend.User
	query := fmt.Sprintf(("SELECT id FROM %s WHERE username = $1 AND password_hash = $2"), userTable)
	err := r.db.Get(&user, query, username,password)
	return user,err
}