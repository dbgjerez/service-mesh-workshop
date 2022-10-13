package interfaces

import (
	"error-service/domain/dto"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	serviceVersion = "SERVICE_VERSION"
	serviceName    = "SERVICE_NAME"
)

type StatusHandler struct {
	status dto.Status
}

func NewStatusHandler() (statusHandler *StatusHandler) {
	app := dto.App{}
	app.Version = os.Getenv(serviceVersion)
	app.Service = os.Getenv(serviceName)
	status := dto.Status{App: app}
	return &StatusHandler{status: status}
}

func (handler *StatusHandler) StatusGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, handler.status)
	}
}
