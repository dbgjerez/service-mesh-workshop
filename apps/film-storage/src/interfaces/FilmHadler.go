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
	ParamPremium = "premium"
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
		f := handler.repository.FindById(id)
		c.JSON(http.StatusOK, f)
	}
}

func (handler *FilmHandler) FilmGetAllHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		premiumValue := c.DefaultQuery("premium", "false")
		premium, err := strconv.ParseBool(premiumValue)
		if err != nil {
			// TODO: exception error
			log.Fatalf("Not able to parse premium value %s", err)
		}
		films := handler.repository.GetAllByType(premium)
		c.JSON(http.StatusOK, films)
	}
}
