package service

import (
	backend "cmd/main.go"
	"cmd/main.go/pkg/repository"
)


type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Administration interface{
	AddBikes(backend.Bike) (int,error)
	GetBikes() ([]backend.Bike)
	UpdateBikes(bike backend.Bike) (error)
	DeleteBikes(bikeId int) (error)
	AddRollers(rollers backend.Rollers) (int,error)
	GetRollers() ([]backend.Rollers)
	UpdateRollers(rollers backend.Rollers) error
	DeleteRollers(rollerId int) error

	
}
type BikeReservation interface {
	CreateBikeReservation(userId int , res backend.BikeReservation) (int ,error)
	GetBikeReservation(userId int) ([]backend.BikeReservation,error)
	UpdateBikeReservation(userId int , res backend.BikeReservation) error
	DeleteBikesReservation(userId, bikeReservId int) error
}
type RollerReservation interface{
	ReservRoller(backend.RollersReservation) (int, error)
	GetReservations(userId int) ([]backend.BikeReservation, []backend.RollersReservation)
	UpdateRollersReserv(backend.RollersReservation) error
	DeleteRollersReserv(rollerReservId int) error
}


type Service struct {
	Authorization
	Administration
	BikeReservation
	RollerReservation
}

func  NewService(repos *repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Administration: NewAdminService(repos.Admin),
		BikeReservation: NewBikeReservationService(repos.BikeReservation),
	}
}