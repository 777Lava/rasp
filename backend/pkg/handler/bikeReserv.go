package handler

import (
	backend "cmd/main.go"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type getReservBike struct {
	Id       int    `json:"id"`
	BikeId   int    `json:"bikeId"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
}

func (h *Handler) createBikeReserv(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		return
	}
	var get getReservBike
	var input backend.BikeReservation
	if err := c.BindJSON(&get); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Парсим строку в объект time.Time
	timecheckin, err := time.Parse("02-01-2006", get.Checkin)
	timechekout, err := time.Parse("02-01-2006", get.Checkout)
	if err != nil {
		fmt.Println("Ошибка при парсинге строки:", err)
		return
	}
	input.BikeId = get.BikeId
	input.UserId = userId
	input.Checkin = timecheckin
	input.Checkout = timechekout
	id, err := h.services.CreateBikeReservation(userId, input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getBikeReserv(c *gin.Context) {
	var resBike []getReservBike

	userId, err := GetUserId(c)
	if err != nil {
		return
	}
	reservs, err := h.services.GetBikeReservation(userId)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	for _, reserv := range reservs {
		resBike = append(resBike, getReservBike{
			Id:       reserv.Id,
			BikeId:   reserv.BikeId,
			Checkin:  reserv.Checkin.Format("02-01-2006"),
			Checkout: reserv.Checkout.Format("02-01-2006"),
		})

	}

	c.JSON(http.StatusOK, resBike)

}

func (h *Handler) updateBikeReserv(c *gin.Context) {

	var get getReservBike
	var input backend.BikeReservation
	userId, err := GetUserId(c)
	if err != nil {
		return
	}
	if err := c.BindJSON(&get); err != nil {
		fmt.Println("badstatus ", err.Error())
		return
	}
	timecheckin, err := time.Parse("02-01-2006", get.Checkin)
	timechekout, err := time.Parse("02-01-2006", get.Checkout)
	input.Id = get.Id
	input.BikeId = get.BikeId
	input.Checkin = timecheckin
	input.Checkout = timechekout
	err = h.services.UpdateBikeReservation(userId, input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) deleteBikeReserv(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print("invalid id param")
		return
	}
	err = h.services.DeleteBikesReservation(userId, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}
