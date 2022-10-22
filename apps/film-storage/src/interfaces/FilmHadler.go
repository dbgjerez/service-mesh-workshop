package interfaces

import (
	"film-storage/domain/dto"
	"film-storage/domain/model"
	"film-storage/service"
	"fmt"
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
			// TODO: exception error
			log.Fatalf("Error al recuperar %s", err)
		}
		c.JSON(http.StatusOK, list)
	}
}

func (handler *FilmHandler) FilmCreateHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		reqFilm := dto.FilmCreateDTO{}
		if err := c.BindJSON(&reqFilm); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Errorf("%s", err))
			return
		}
		f := service.FilmCreateDTOToFilm(reqFilm)
		f, err := handler.repository.CreateOne(f)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Errorf("%s", err))
			return
		}
		c.JSON(http.StatusCreated, f)
	}
}