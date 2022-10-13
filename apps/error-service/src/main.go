package main

import (
	"error-service/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")
	{
		h := interfaces.HealthcheckHandler{}
		v1.GET("/health", h.HealthcheckGetHandler())

		s := interfaces.NewInfoHandler()
		v1.GET("/info", s.InfoGetHandler())
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}
