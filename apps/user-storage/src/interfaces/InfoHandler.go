package interfaces

import (
	"net/http"
	"user-storage/domain/dto"
	"user-storage/utils"

	"github.com/gin-gonic/gin"
)

const (
	serviceVersion = "SERVICE_VERSION"
	serviceName    = "SERVICE_NAME"
	buildTime      = "SERVICE_BUILD_TIME"
)

type InfoHandler struct {
	info *dto.Info
}

func NewInfoHandler() (infoHandler *InfoHandler) {
	app := dto.App{}
	app.Version = utils.GetEnv(serviceVersion, "")
	app.Service = utils.GetEnv(serviceName, "")
	app.BuildTime = utils.GetEnv(buildTime, "")
	if app.Version == "" || app.Service == "" {
		//FIXME: añadir error controlado
	}
	info := dto.Info{App: app}
	return &InfoHandler{info: &info}
}

func (handler *InfoHandler) InfoGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, handler.info)
	}
}
