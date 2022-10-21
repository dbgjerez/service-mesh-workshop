package interfaces

import (
	"film-storage/domain/dto"
	"film-storage/domain/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
	filmDao *model.FilmRepository
}

func NewHealthcheckHandler(dao *model.FilmRepository) *HealthcheckHandler {
	return &HealthcheckHandler{filmDao: dao}
}

func (handler *HealthcheckHandler) HealthcheckGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		h := dto.Health{}
		err := handler.filmDao.HealthCheck()
		if err != nil {
			log.Printf("%s", err)
			h.Status = dto.HealhStatusDown
		} else {
			h.Status = dto.HealhStatusUp
		}
		c.JSON(http.StatusOK, h)
	}
}
