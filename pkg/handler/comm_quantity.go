package handler

import (
	"github.com/Zavr22/courseGo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type reqOfferBody struct {
	Offer courseGo.CommQuantity `json:"offer"`
}

func (h *Handler) approveQ(c *gin.Context) {
	var offer reqOfferBody
	userIdStr := c.GetHeader("Authorization")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if err := c.BindJSON(&offer); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.ApproveQuantity(userId, offer.Offer.Id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "approved")
}

type getAllCommOResponse struct {
	Data []courseGo.CommQuantity `json:"data"`
}

func (h *Handler) getAllCommO(c *gin.Context) {
	commO, err := h.services.MakeQuantity.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommOResponse{
		Data: commO,
	})
}
