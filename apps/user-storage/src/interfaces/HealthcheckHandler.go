package interfaces

import (
	"log"
	"net/http"
	"user-storage/domain/dto"
	"user-storage/domain/model"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
	userDao *model.UserRepository
}

func NewHealthcheckHandler(dao *model.UserRepository) *HealthcheckHandler {
	return &HealthcheckHandler{userDao: dao}
}

func (handler *HealthcheckHandler) HealthcheckGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		h := dto.Health{}
		err := handler.userDao.HealthCheck()
		if err != nil {
			log.Printf("%s", err)
			h.Status = dto.HealhStatusDown
		} else {
			h.Status = dto.HealhStatusUp
		}
		c.JSON(http.StatusOK, h)
	}
}
