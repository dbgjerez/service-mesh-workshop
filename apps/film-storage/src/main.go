package main

import (
	"context"
	"film-storage/domain/model"
	"film-storage/interfaces"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	filmRepository, err := model.NewFilmRepository(store)
	if err != nil {
		log.Fatalf("Fails reading store %v", err)
	}

	v1 := router.Group("/api/v1")
	{
		h := interfaces.NewHealthcheckHandler(filmRepository)
		v1.GET("/health", h.HealthcheckGetHandler())
	}

	film := router.Group("/api/v1/film")
	{
		f := interfaces.NewFilmHandler(filmRepository)
		film.GET("", f.FilmGetAllHandler())
		film.GET("/:id", f.FilmFindByIdHandler())
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
