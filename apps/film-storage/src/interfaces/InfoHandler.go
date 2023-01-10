package interfaces

import (
	"film-storage/domain/model"
	"film-storage/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	serviceVersion = "SERVICE_VERSION"
	serviceName    = "SERVICE_NAME"
)

type InfoHandler struct {
	info *model.Info
}

func NewInfoHandler() (infoHandler *InfoHandler) {
	app := model.App{}
	app.Version = utils.GetEnv(serviceVersion, "")
	app.Service = utils.GetEnv(serviceName, "")
	if app.Version == "" || app.Service == "" {
		//FIXME: añadir error controlado
	}
	info := model.Info{App: app}
	return &InfoHandler{info: &info}
}

func (handler *InfoHandler) InfoGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, handler.info)
	}
}
