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

func (handler *FilmHandler) FilmFindByIdHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		filmId := c.Param(ParamIDName)
		log.Print("id: ", filmId)
		f, err := handler.repository.FindById(filmId)
		if err != nil {
			// TODO: exception error
			log.Fatalf("Error al recuperar %s", err)
		}
		film := dto.FilmDTO{
			Id:       f.Id.Hex(),
			Duration: int(f.Duration),
			Premium:  f.Premium,
			Title:    f.Title,
		}
		c.JSON(http.StatusOK, film)
	}
}

func (handler *FilmHandler) FilmGetAllHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		list, err := handler.repository.GetAll()
		if err != nil {
			// TODO: exception error
			log.Fatalf("Error al recuperar %s", err)
		}
		res := []dto.FilmDTO{}
		for _, f := range list {
			res = append(res, dto.FilmDTO{
				Id:       f.Id.Hex(),
				Duration: int(f.Duration),
				Premium:  f.Premium,
				Title:    f.Title,
			})
		}
		c.JSON(http.StatusOK, res)
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
