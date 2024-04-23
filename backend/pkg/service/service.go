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
type Reservation interface {
	ReservBike(backend.BikeReservation) (int ,error)
	ReservRoller(backend.RollersReservation) (int, error)
}
type Service struct {
	Authorization
	Administration
}

func  NewService(repos *repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Administration: NewAdminService(repos.Admin),
	}
}