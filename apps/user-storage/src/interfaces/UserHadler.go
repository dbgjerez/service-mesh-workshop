package interfaces

import (
	"net/http"
	"user-storage/domain/dto"
	"user-storage/domain/model"

	"github.com/gin-gonic/gin"
)

const (
	ParamIDName = "username" // find by id
)

type UserHandler struct {
	repository *model.UserRepository
}

func NewFilmHandler(dao *model.UserRepository) *UserHandler {
	return &UserHandler{repository: dao}
}

func (handler *UserHandler) FilmFindByIdHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.Param(ParamIDName)
		f, err := handler.repository.FindById(username)
		if err != nil {
			if err == model.UserNotFound {
				c.JSON(http.StatusNotFound, dto.ErrorDTO{
					Code:    404,
					Message: model.UserNotFound.Error(),
				})
			}
		} else {
			c.JSON(http.StatusOK, f)
		}
	}
}

func (handler *UserHandler) FilmGetAllHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		films := handler.repository.GetAll()
		c.JSON(http.StatusOK, films)
	}
}
