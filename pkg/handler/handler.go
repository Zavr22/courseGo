package handler

import (
	"github.com/Zavr22/courseGo/pkg/service"
	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	allowedHeaders := []string{"Content-Type", "Authorization", "X-Requested-With", "X-User-Id"}
	config.AllowHeaders = allowedHeaders
	router.Use(cors.New(config))
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
		prod.POST("/sortPrByPrice", h.sortProjByPrice)
		prod.POST("/sortVWByPrice", h.sortVWByPrice)
		prod.POST("/sortMonByPrice", h.sortMonByPrice)

	}
	commO := router.Group("/commO")
	{
		commO.GET("/commO", h.getAllCommO)
		commO.POST("/confirm", h.approveQ)
		commO.GET("/getAllForMng", h.getAllCommOMng)
	}
	settings := router.Group("/settings")
	{
		settings.POST("/profit", h.setProfit)
	}

	return router
}
