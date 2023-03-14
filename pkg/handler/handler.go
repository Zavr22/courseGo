package handler

import (
	"github.com/Zavr22/courseGo/pkg/service"
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
		auth.POST("/signUp", h.signUp)
		auth.POST("/signIn", h.signIn)
	}

	prod := router.Group("/prod")
	{
		prod.GET("/proj", h.getAllProjectors)
		prod.GET("/videoW", h.getAllVideoWalls)
		prod.GET("/monitors", h.getAllMonitors)
		prod.GET("/mounts", h.getAllMounts)
		prod.POST("/getPrE", h.getPrWithExtra)
		prod.POST("/confirm", h.confirmOffer)

	}

	return router
}
