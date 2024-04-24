package service

import (
	backend "cmd/main.go"
	"cmd/main.go/pkg/repository"
)






type BikeReservationService struct {
	repo repository.BikeReservation
}

func NewBikeReservationService(repo repository.BikeReservation) *BikeReservationService {
	return &BikeReservationService{repo: repo}
}


func (r *BikeReservationService) CreateBikeReservation(userId int, res backend.BikeReservation) (int ,error){
	return r.repo.CreateBikeReservation(userId,res)	
}

func (r *BikeReservationService) GetBikeReservation(userId int) ([]backend.BikeReservation,error){
	return r.repo.GetBikeReservation(userId)
}
func (r *BikeReservationService) UpdateBikeReservation(userId int , res backend.BikeReservation) error{
	return r.repo.UpdateBikeReservation(userId, res)
}
func (r *BikeReservationService)DeleteBikesReservation(userId,bikeReservId int) error{
	return	 r.repo.DeleteBikesReservation(userId,bikeReservId)
}
