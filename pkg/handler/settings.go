package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type settingsRequest struct {
	Roi string `json:"roi"`
}

func (h *Handler) setProfit(c *gin.Context) {
	var roiParam settingsRequest
	userIdStr := c.GetHeader("X-User-Id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if userId != 1 {
		c.JSON(http.StatusForbidden, "you are not a director")
	}
	if err := c.BindJSON(&roiParam); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	roiFloat, err := strconv.ParseFloat(roiParam.Roi, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.SetProfit(roiFloat); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Settings was set")
}
