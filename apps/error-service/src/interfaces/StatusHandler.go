package interfaces

import (
	"github.com/gin-gonic/gin"
)

type StatusHandler struct {
}

func (handler *StatusHandler) StatusGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}
