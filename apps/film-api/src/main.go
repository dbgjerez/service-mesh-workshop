package main

import (
	"film-api/domain/repository"
	"film-api/interfaces"
	"film-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	SERVER_PORT string = ":8000"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.Default())

	fService := services.NewFilmService(
		repository.NewFilmRepository(),
		repository.NewUserRepository(),
		repository.NewCommentRepository(),
	)
	tService := services.NewTokenService()

	v1 := router.Group("/api/v1")
	{
		h := interfaces.NewHealthcheckHandler()
		v1.GET("/health", h.HealthcheckGetHandler())

		s := interfaces.NewInfoHandler()
		v1.GET("/info", s.InfoGetHandler())

		f := interfaces.NewFilmHandler(fService, tService)
		v1.GET("/films", f.GetFilmsByUserHandler())
	}

	router.Run(SERVER_PORT)
}
