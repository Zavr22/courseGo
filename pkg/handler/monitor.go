package handler

import (
	"github.com/Zavr22/courseGo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllMonitorsResponse struct {
	Data []courseGo.Monitor `json:"data"`
}

func (h *Handler) getAllMonitors(c *gin.Context) {
	monitors, err := h.services.Monitor.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllMonitorsResponse{
		Data: monitors,
	})
}
