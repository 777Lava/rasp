package repository

import (
	backend "cmd/main.go"

	"github.com/jmoiron/sqlx"
)

type bikeReservationPostgres struct{
	db *sqlx.DB
}
func NewBikeReservPostgres(db *sqlx.DB) *bikeReservationPostgres {
	return &bikeReservationPostgres{db: db}
}

	
	
	
func (*bikeReservationPostgres) CreateBikeReservation(backend.BikeReservation) (int, error){
	return 1,nil
}

func (*bikeReservationPostgres) GetBikeReservation(userId int) ([]backend.BikeReservation){
	return nil
}
func (*bikeReservationPostgres) DeleteBikesReservation(bikeReservId int) (error){
	return nil
}
func (*bikeReservationPostgres) UpdateBikeReservation(backend.BikeReservation) (error){
	return	 nil
}