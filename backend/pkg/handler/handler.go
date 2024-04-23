package handler

import (
	"cmd/main.go/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{

		bikes := api.Group("/bikes")
		{
			bikes.POST("/")

		}
		rollers := api.Group("/rollers")
		{
			rollers.POST("/")
		}
		api.GET("/")
	}
	admin := router.Group("/admin", h.adminIdentity)
	{
		bikes := admin.Group("/bikes")
		{
			bikes.POST("/", h.addBikes)
			bikes.GET("/", h.getBikes)
			bikes.PUT("/",h.updateBikes)
			bikes.DELETE("/:id", h.deleteBikes)
		}
		rollers := admin.Group("/rollers")
		{
			rollers.POST("/",h.addRollers)
			rollers.GET("/",h.getRollers)
			rollers.PUT("/", h.updateRollers)
			rollers.DELETE("/:id",h.deleteRollers)
		}
	}

	return router
}
