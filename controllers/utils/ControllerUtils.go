package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pollp/models"
)

func HandleControllerError(err error, c *gin.Context) bool {
	if err != nil {
		var httpError *models.HttpError
		switch {
		case errors.As(err, &httpError):
			c.AbortWithStatus(httpError.HttpStatusCode)
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	return err != nil
}
