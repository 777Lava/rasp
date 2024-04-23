package main

import (
	backend "cmd/main.go"
	"cmd/main.go/pkg/handler"
	"cmd/main.go/pkg/repository"
	"cmd/main.go/pkg/service"
	"fmt"


	_ "github.com/lib/pq"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		fmt.Println("ошибка в подключении кбд")

		fmt.Println(err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(backend.Server)
	srv.Run("8000", handlers.InitRoutes())
}
