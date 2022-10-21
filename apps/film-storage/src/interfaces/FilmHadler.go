package interfaces

import (
	"film-storage/domain/model"
	"log"
	"net/http"

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

func (handler *FilmHandler) FilmGetAllHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		list, err := handler.repository.GetAll()
		if err != nil {
			log.Fatalf("Error al recuperar %s", err)
		}
		c.JSON(http.StatusOK, list)
	}
}
