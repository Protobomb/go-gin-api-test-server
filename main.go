package main

import (
	"net/http"

	"github.com/Protobomb/squidge/squidge"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("/opt/go-gin-api-test-server/templates/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Squidge Website",
		})
	})
	router.GET("/squidge", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": squidge.Squidge(),
		})
	})
	router.Run(":8080")
}
