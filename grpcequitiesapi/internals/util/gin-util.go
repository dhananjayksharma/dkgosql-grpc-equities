package util

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: These are optional to use. Create more here if you need extra utilities for working with Gin.

// successResponse send successful response back to client.
func successResponse(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func HandleError(c *gin.Context, err error) {
	var errBadRequest *BadRequest
	var errNotFound *NotFound
	var errDataNotFound *DataNotFound
	var errUnauthorized *UnAuthorized
	var errInternalServer *InternalServer

	var statusCode int
	var errMessage string
	switch true {
	case errors.As(err, &errBadRequest):
		statusCode = http.StatusBadRequest
		errMessage = errBadRequest.ErrMessage
	case errors.As(err, &errNotFound):
		statusCode = http.StatusNotFound
		errMessage = errNotFound.ErrMessage
	case errors.As(err, &errDataNotFound):
		statusCode = http.StatusOK
		errMessage = errDataNotFound.ErrMessage
	case errors.As(err, &errUnauthorized):
		statusCode = http.StatusUnauthorized
		errMessage = errUnauthorized.ErrMessage
	case errors.As(err, &errInternalServer):
		statusCode = http.StatusInternalServerError
		errMessage = errInternalServer.ErrMessage
	default:
		statusCode = http.StatusInternalServerError
		errMessage = err.Error()
	}

	errorResponse(c, statusCode, errMessage)
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	errorBody := gin.H{
		"error":   true,
		"status":  statusCode,
		"message": message,
	}

	c.JSON(statusCode, errorBody)
}
