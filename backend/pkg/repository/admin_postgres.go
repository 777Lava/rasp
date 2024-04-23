package repository

import (
	backend "cmd/main.go"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Administration struct {
	db *sqlx.DB
}

func NewAdministratrion(db *sqlx.DB) *Administration {
	return &Administration{db: db}
}

func (a *Administration) AddBikes(bike backend.Bike) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createBikeQuery := fmt.Sprintf("insert into %s (name, price, description) values ($1, $2, $3) returning id", bikesTable)
	row := tx.QueryRow(createBikeQuery, bike.Name, bike.Price, bike.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (a *Administration) GetBikes() ([]backend.Bike, error) {
	var bikes []backend.Bike
	query := fmt.Sprintf(`SELECT * FROM %s`, bikesTable)
	err := a.db.Select(&bikes, query)
	return bikes, err
}
func (a *Administration) UpdateBikes(bike backend.Bike) error {
	query := fmt.Sprintf(`update %s bi set name = $2, price=$3,description=$4 where bi.id = $1`, bikesTable)
	_, err := a.db.Exec(query, bike.Id, bike.Name, bike.Price, bike.Description)
	return err
}
func (a *Administration) DeleteBikes(bikeId int) error {
	query := fmt.Sprintf(`delete from %s tl where tl.id = $1`, bikesTable)
	_, err := a.db.Exec(query, bikeId)
	return err
}

func (a *Administration) AddRollers(rollers backend.Rollers) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createBikeQuery := fmt.Sprintf("insert into %s (name, price, description, size) values ($1, $2, $3, $4) returning id", rollersTable)
	row := tx.QueryRow(createBikeQuery, rollers.Name, rollers.Price, rollers.Description, rollers.Size)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}
func (a *Administration) GetRollers() ([]backend.Rollers, error) {
	var rollers []backend.Rollers
	query := fmt.Sprintf(`SELECT * FROM %s`, rollersTable)
	err := a.db.Select(&rollers, query)
	return rollers, err
}
func (a *Administration) UpdateRollers(rollers backend.Rollers) error {
	query := fmt.Sprintf(`update %s bi set name = $2, price=$3,description=$4,size =$5 where bi.id = $1`, rollersTable)
	_, err := a.db.Exec(query, rollers.Id, rollers.Name, rollers.Price, rollers.Description, rollers.Size)
	return err
}
func (a *Administration) DeleteRollers(rollerId int) error {
	query := fmt.Sprintf(`delete from %s tl where tl.id = $1`, rollersTable)
	_, err := a.db.Exec(query, rollerId)
	return err
}
