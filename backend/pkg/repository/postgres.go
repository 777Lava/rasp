package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable               = "users"
	bikeReservationTable    = "bikeReservation"
	bikesTable              = "bikes"
	rollersTable            = "rollers"
	rollersReservationTable = "rollersReservation"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ошибка в пинге")
		return nil, err
	}

	return db, nil
}
