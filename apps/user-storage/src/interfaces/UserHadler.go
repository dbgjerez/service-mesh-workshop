package interfaces

import (
	"log"
	"net/http"
	"strconv"
	"user-storage/domain/model"

	"github.com/gin-gonic/gin"
)

const (
	ParamIDName = "id" // find by id
)

type UserHandler struct {
	repository *model.UserRepository
}

func NewFilmHandler(dao *model.UserRepository) *UserHandler {
	return &UserHandler{repository: dao}
}

func (handler *UserHandler) FilmFindByIdHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		filmId := c.Param(ParamIDName)
		id, err := strconv.Atoi(filmId)
		if err != nil {
			// TODO: exception error
			log.Fatalf("Not able to parse the id %s", err)
		}
		f := handler.repository.FindById(id)
		c.JSON(http.StatusOK, f)
	}
}

func (handler *UserHandler) FilmGetAllHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		films := handler.repository.GetAll()
		c.JSON(http.StatusOK, films)
	}
}
