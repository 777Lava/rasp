package repository

import (


	"github.com/jmoiron/sqlx"
)

type rollerReservation struct {
	db sqlx.DB
}

func NewRollerReservation(db sqlx.DB) rollerReservation { return rollerReservation{db: db} }