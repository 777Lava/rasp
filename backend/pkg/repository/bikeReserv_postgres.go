package repository

import (
	backend "cmd/main.go"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type bikeReservationPostgres struct {
	db *sqlx.DB
}

func NewBikeReservPostgres(db *sqlx.DB) *bikeReservationPostgres {
	return &bikeReservationPostgres{db: db}
}

type getTimePostgres struct {
	Id       int   `json:"id" db:"id"`
	Checkin  int64 `json:"checkin" db:"checkin"`
	Checkout int64 `json:"checkout" db:"checkout"`
}
type getBikePostgres struct {
	Id       int   `json:"id" db:"id"`
	UserId   int   `json:"userId db:"user_id"`
	BikeId   int   `json:"bikeId" db:"bike_id"`
	Checkin  int64 `json:"checkin" db:"checkin"`
	Checkout int64 `json:"checkout" db:"checkout"`
}

func (a *bikeReservationPostgres) CreateBikeReservation(userId int, res backend.BikeReservation) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	var lists []getTimePostgres
	checktimeQuery := fmt.Sprintf(`select id, checkin, checkout from %s where bike_id = $1`, bikeReservationTable)
	err = a.db.Select(&lists, checktimeQuery, res.BikeId)
	for _, list := range lists {
		if (res.Checkin.Unix()-list.Checkin >= 0 && (res.Checkin.Unix()-list.Checkout <= 0)) ||
			(res.Checkout.Unix()-list.Checkin >= 0 && res.Checkout.Unix()-list.Checkout <= 0) {
			return 0, fmt.Errorf("недопустимая дата")
		}
	}

	createBikerResQuery := fmt.Sprintf("insert into %s (user_id, bike_id, checkin, checkout) values ($1, $2, $3,$4) returning id", bikeReservationTable)
	row := tx.QueryRow(createBikerResQuery, res.UserId, res.BikeId, res.Checkin.Unix(), res.Checkout.Unix())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (a *bikeReservationPostgres) GetBikeReservation(userId int) ([]backend.BikeReservation, error) {
	var get []getBikePostgres
	var lists []backend.BikeReservation

	query := fmt.Sprintf(`SELECT id, bike_id, checkin, checkout FROM %s where user_id = $1`, bikeReservationTable)

	err := a.db.Select(&get, query, userId)
	if err != nil {
		fmt.Println("ошибка в запросе", err.Error())
		return nil, err
	}

	for _, list := range get {
		lists = append(lists, backend.BikeReservation{
			Id:       list.Id,
			UserId:   userId,
			BikeId:   list.BikeId,
			Checkin:  time.Unix(list.Checkin, 0),
			Checkout: time.Unix(list.Checkout, 0),
		})

	}

	return lists, err
}

func (a *bikeReservationPostgres) UpdateBikeReservation(userId int, res backend.BikeReservation) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	var lists []getTimePostgres
	checktimeQuery := fmt.Sprintf(`select id, checkin, checkout from %s where bike_id=$1`, bikeReservationTable)
	err = a.db.Select(&lists, checktimeQuery, res.BikeId)
	if err != nil {
		fmt.Println("ошибка в селекте ", err.Error())
		return err

	}
	for _, list := range lists {
		if (res.Checkin.Unix()-list.Checkin >= 0 && (res.Checkin.Unix()-list.Checkout <= 0)) ||
			(res.Checkout.Unix()-list.Checkin >= 0 && res.Checkout.Unix()-list.Checkout <= 0) {
			return fmt.Errorf("недопустимая дата")
		}
	}
	createBikerResQuery := fmt.Sprintf(`update %s set bike_id=$2, checkin=$3, checkout=$4 where user_id=$1 and id=$5`, bikeReservationTable)
	_, err = a.db.Exec(createBikerResQuery, userId, res.BikeId, res.Checkin.Unix(), res.Checkout.Unix(), res.Id)
	if err != nil {
		fmt.Println("ошибка в запросе", err.Error())
		return err
	}

	return tx.Commit()
}

func (a *bikeReservationPostgres) DeleteBikesReservation(userId, bikeReservId int) error {
	query := fmt.Sprintf(`delete from %s tl where user_id = $1 and tl.id = $2`, bikeReservationTable)
	_, err := a.db.Exec(query, userId, bikeReservId)
	return err
}
