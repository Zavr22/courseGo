package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type settingsRequest struct {
	roi string `json:"roi"`
}

func (h *Handler) setProfit(c *gin.Context) {
	var roiParam settingsRequest
	userIdStr := c.GetHeader("Authorization")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if userId != 1 {
		c.JSON(http.StatusOK, "you are not a director")
	}
	if err := c.BindJSON(&roiParam); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	roiFloat, err := strconv.ParseFloat(roiParam.roi, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.SetProfit(userId, roiFloat); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Settings was set")
}
