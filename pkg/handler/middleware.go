package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userCtx = "userId"
)

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is  not of valid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
