package repository

import (
	backend "cmd/main.go"


	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GetUser(username, password string) (backend.User, error)
}

type BikeReservation interface {
	CreateBikeReservation(backend.BikeReservation) (int, error)
	GetBikeReservation(userId int) ([]backend.BikeReservation)
	DeleteBikesReservation(bikeReservId int) (error)
	UpdateBikeReservation(backend.BikeReservation) (error)
}
type RollersReservation interface{
	CreateRollersReservation(backend.RollersReservation) (int, error)
	GetRollerReservation(userId int) ([]backend.RollersReservation)
	UpdateRollerReservation(backend.RollersReservation) (error)
	DeleteRollersReservation(rollerReservId int) (error)
}

type Admin interface{
	AddBikes(bike backend.Bike) (int,error)
	GetBikes() ([]backend.Bike,error)
	UpdateBikes(bike backend.Bike) (error)
	DeleteBikes(bikeId int) (error)
	AddRollers(rollers backend.Rollers) (int,error)
	GetRollers() ([]backend.Rollers,error)
	UpdateRollers(rollers backend.Rollers) error
	DeleteRollers(rollerId int) error
}


type Repository struct {
	Authorization
	BikeReservation
	RollersReservation
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BikeReservation: NewBikeReservPostgres(db),
		Admin : NewAdministratrion(db),
		
	}
}