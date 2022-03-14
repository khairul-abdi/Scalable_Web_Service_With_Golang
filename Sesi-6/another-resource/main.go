package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/response-use-html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"judul": "Response dengan Output HTML",
		})
	})

	router.GET("/response-use-string", func(c *gin.Context) {
		c.String(http.StatusOK, "Cara menggunakan Output String Gin")
	})

	router.GET("/response-use-json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Ini Pesan JSON", "status": http.StatusOK})
	})

	router.Run(":8080")
}
