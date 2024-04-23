package backend

import "time"



type Bike struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
	Price int`json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}

type Rollers struct{
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Price int `json:"price" db:"price" binding:"required"`
	Description string `json:"description" db:"description"`
	Size float64 `json:"size" db:"size" binding:"required"`
}

type BikeReservation struct {
	UserId int 
	BikeId int
	Start time.Time
	Finish time.Time 
	
}

type RollersReservation struct {
	UserId int 
	RollersId int
	Start time.Time
	Finish time.Time 
}