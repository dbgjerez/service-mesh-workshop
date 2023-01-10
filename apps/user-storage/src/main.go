package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user-storage/domain/model"
	"user-storage/interfaces"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.Default())

	store := model.NewStore()
	filmRepository, err := model.NewUserRepository(store)
	if err != nil {
		log.Fatalf("Fails reading store %v", err)
	}

	v1 := router.Group("/api/v1")
	{
		h := interfaces.NewHealthcheckHandler(filmRepository)
		v1.GET("/health", h.HealthcheckGetHandler())

		s := interfaces.NewInfoHandler()
		v1.GET("/info", s.InfoGetHandler())
	}

	film := router.Group("/api/v1/user")
	{
		f := interfaces.NewFilmHandler(filmRepository)
		film.GET("", f.FilmGetAllHandler())
		film.GET("/:username", f.FilmFindByIdHandler())
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	srv.Shutdown(ctx)
	os.Exit(0)

}
