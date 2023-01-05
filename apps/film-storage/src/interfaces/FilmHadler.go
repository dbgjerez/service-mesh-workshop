package interfaces

import (
	"film-storage/domain/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	ParamIDName  = "id" // find by id
	BodyDataName = "data"
	BodyDataSize = "size"
)

type FilmHandler struct {
	repository *model.FilmRepository
}

func NewFilmHandler(dao *model.FilmRepository) *FilmHandler {
	return &FilmHandler{repository: dao}
}

func (handler *FilmHandler) FilmFindByIdHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		filmId := c.Param(ParamIDName)
		id, err := strconv.Atoi(filmId)
		if err != nil {
			// TODO: exception error
			log.Fatalf("Not able to parse the id %s", err)
		}
		f, err := handler.repository.FindById(id)
		if err != nil {
			// TODO: exception error
			log.Fatalf("Error al recuperar %s", err)
		}
		c.JSON(http.StatusOK, f)
	}
}

func (handler *FilmHandler) FilmGetAllHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		films, err := handler.repository.GetAll()
		if err != nil {
			// TODO: exception error
			log.Fatalf("Error al recuperar %s", err)
		}
		c.JSON(http.StatusOK, films)
	}
}
