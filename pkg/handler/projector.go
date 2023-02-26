package handler

import (
	"github.com/Zavr22/courseGo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllProjectorsResponse struct {
	Data []courseGo.Projector `json:"data"`
}

func (h *Handler) getAllProjectors(c *gin.Context) {
	projectors, err := h.services.Projector.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllProjectorsResponse{
		Data: projectors,
	})
}
