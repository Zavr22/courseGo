package handler

import (
	"github.com/Zavr22/courseGo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Products struct {
	monitor   courseGo.Monitor
	projector courseGo.Projector
	videoWall courseGo.VideoWall
	mount     courseGo.Mount
}

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

type getAllVideoWallResponse struct {
	Data []courseGo.VideoWall `json:"data"`
}

func (h *Handler) getAllVideoWalls(c *gin.Context) {
	videoWalls, err := h.services.VideoWall.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllVideoWallResponse{
		Data: videoWalls,
	})
}

type getAllMountsResponse struct {
	Data []courseGo.Mount `json:"data"`
}

func (h *Handler) getAllMounts(c *gin.Context) {
	mounts, err := h.services.Mount.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllMountsResponse{
		Data: mounts,
	})
}

type getProjWithExtraResp struct {
	Data []courseGo.ProdInventory `json:"data"`
}

type getMonWithExtraResp struct {
	Data []courseGo.ProdInventory `json:"data"`
}

type getVideoWallWithExtraResp struct {
	Data []courseGo.ProdInventory `json:"data"`
}

func (h *Handler) getPrWithExtra(c *gin.Context) {
	var input courseGo.Params
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	switch input.CategoryId {
	case 1:
		prod, err := h.services.PickUpProjectorWithExtra(input)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, getProjWithExtraResp{
			Data: prod,
		})
	case 2:
		prod, err := h.services.PickUpMonitorWithExtra(input)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, getMonWithExtraResp{
			Data: prod,
		})
	case 3:
		prod, err := h.services.PickUpVideoWallWithExtra(input)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, getVideoWallWithExtraResp{
			Data: prod,
		})

	}

}

func (h *Handler) confirmOffer() {

}