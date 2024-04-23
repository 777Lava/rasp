package service

import (
	backend "cmd/main.go"
	"cmd/main.go/pkg/repository"
	"fmt"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (a *AdminService) AddBikes(bike backend.Bike) (int, error) {
	return a.repo.AddBikes(bike)
}
func (a *AdminService) GetBikes() []backend.Bike {
	bikes, err := a.repo.GetBikes()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return bikes
}
func (a *AdminService) UpdateBikes(bike backend.Bike) error {
	return a.repo.UpdateBikes(bike)
}
func (a *AdminService) DeleteBikes(bikeId int) error {
	return a.repo.DeleteBikes(bikeId)
}

func (a *AdminService) AddRollers(rollers backend.Rollers) (int, error) {
	return a.repo.AddRollers(rollers)
}
func (a *AdminService) GetRollers() []backend.Rollers {
	rollers,err := a.repo.GetRollers()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return rollers
}
func (a *AdminService) UpdateRollers(rollers backend.Rollers) error {
	return a.repo.UpdateRollers(rollers)
}
func (a *AdminService) DeleteRollers(rollerId int) error {
	return a.repo.DeleteRollers(rollerId)
}
