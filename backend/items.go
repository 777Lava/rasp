package backend

import "time"

type Bike struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}

type Rollers struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Price       int     `json:"price" db:"price" binding:"required"`
	Description string  `json:"description" db:"description"`
	Size        float64 `json:"size" db:"size" binding:"required"`
}

type BikeReservation struct {
	Id       int       `json:"-" db:"id"`
	UserId   int       `json:"user_id" db:"user_id"`
	BikeId   int       `json:"bike_id" db:"bike_id"`
	Checkin  time.Time `json:"checkin" db:"checkin"`
	Checkout time.Time `json:"checkout" db:"checkout"`
}

type RollersReservation struct {
	Id        int
	UserId    int
	RollersId int
	Checkin   time.Time
	Checkout  time.Time
}
