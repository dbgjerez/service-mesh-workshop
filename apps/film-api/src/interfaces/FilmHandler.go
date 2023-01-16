package interfaces

import (
	"film-api/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	TOKEN_USER = "user"
)

type FilmHandler struct {
	fService *services.FilmService
	tService *services.TokenService
}

func NewFilmHandler(fService *services.FilmService, tService *services.TokenService) *FilmHandler {
	return &FilmHandler{
		fService: fService,
		tService: tService,
	}
}

func (handler *FilmHandler) GetFilmsByUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")
		username, err := handler.tService.GetClaim(jwtToken, TOKEN_USER)
		if err != nil {
			c.JSON(http.StatusForbidden, fmt.Sprintf("%v", err))
		} else {
			films, err := handler.fService.FindFilmsByUser(username)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			} else {
				c.JSON(http.StatusOK, films)
			}
		}
	}
}
