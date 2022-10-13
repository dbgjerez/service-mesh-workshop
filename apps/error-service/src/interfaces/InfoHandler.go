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

type InfoHandler struct {
	info *dto.Info
}

func NewInfoHandler() (infoHandler *InfoHandler) {
	app := dto.App{}
	app.Version = os.Getenv(serviceVersion)
	app.Service = os.Getenv(serviceName)
	info := dto.Info{App: app}
	return &InfoHandler{info: &info}
}

func (handler *InfoHandler) InfoGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, handler.info)
	}
}
