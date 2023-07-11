package handlers

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler interface {
	Health(c *gin.Context)
}

type healthHandler struct {
}

func NewHealthHandler() HealthHandler {
	return healthHandler{}
}

func (health healthHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{"action": "Health", "status": "success", "message": "Healtch Check OK"})
}
