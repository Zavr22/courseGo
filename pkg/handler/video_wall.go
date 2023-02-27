package handler

import (
	"github.com/Zavr22/courseGo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllVideoWallResponse struct {
	Data []courseGo.VideoWall `json:"data"`
}

func (h *Handler) getAllVideoWalls(c *gin.Context) {
	videoWalls, err := h.services.Projector.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllProjectorsResponse{
		Data: videoWalls,
	})
}
