package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON sends StatusOK, formatted message and data.
//
// Example 1: ginutil.JSON(c, &userInfo, "login success")
// Example 2: ginutil.JSON(c, &userDetails, "Email has been sent to: %s", email)
// Example 3: ginutil.JSON(c, &details, "Email sent to: %s, token will expire in %d minutes", email, expires)
func JSON(c *gin.Context, data interface{}, format string, v ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf(format, v...),
		"error":   false,
		"data":    data,
	})
}

// JSONError sends passed in http status code, formatted message and data.
//
// Example 1: ginutil.JSONError(c, http.StatusNotFound, &userInfo, "login failed, no such user")
func JSONError(c *gin.Context, status int, data interface{}, format string, v ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": fmt.Sprintf(format, v...),
		"error":   true,
		"data":    data,
	})
}
