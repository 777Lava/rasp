package repository

import "github.com/jmoiron/sqlx"

type CheckReservationPostgres struct {
	db *sqlx.DB
}

func NewCheckReservationPostgres(db *sqlx.DB) *CheckReservationPostgres { return &CheckReservationPostgres{db: db}}

func (c *CheckReservationPostgres) CheckBikeReservation(bikeId int) (bool,error){return true,nil}
func (c *CheckReservationPostgres) CheckRollerReservation(rollerId int) (bool,error) {return true,nil}