package handler

import (
	"fmt"
	"net/http"
	"strconv"

	backend "cmd/main.go"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addBikes(c *gin.Context) {

	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	var input backend.Bike

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := h.services.Administration.AddBikes(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getBikes(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	bikes := h.services.GetBikes()
	c.JSON(http.StatusOK, bikes)
}

func (h *Handler) updateBikes(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	var input backend.Bike
	if err := c.BindJSON(&input); err != nil {
		fmt.Println("badstatus ", err.Error())
		return
	}
	h.services.UpdateBikes(input)
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) deleteBikes(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print("invalid id param")
		return
	}
	h.services.DeleteBikes(id)
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) addRollers(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	var input backend.Rollers

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := h.services.Administration.AddRollers(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getRollers(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	rollers := h.services.GetRollers()
	c.JSON(http.StatusOK, rollers)
}

func (h *Handler) updateRollers(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	var input backend.Rollers
	if err := c.BindJSON(&input); err != nil {
		fmt.Println("badstatus ", err.Error())
		return
	}
	err := h.services.UpdateRollers(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) deleteRollers(c *gin.Context) {
	if check := checkAdmin(c); check != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, "not authorized")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	if err != nil {
		fmt.Print("invalid id param")
		return
	}
	h.services.DeleteRollers(id)
	c.JSON(http.StatusOK, "ok")
}
