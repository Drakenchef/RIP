package api

import (
	backend "github.com/drakenchef/RIP"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func StartServer() {
	r := gin.Default()

	r.Static("/resources", "./resources")
	r.Static("/static", "./static")

	r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/index.html")))

	AMS := []backend.AMS{
		{1, "Horizon AMS", "Описание товара 1", "horizon.jpg", "test desc", 200, 4.2},
		{2, "Nasa AMS", "Описание товара 2", "nasa.jpg", "test desc2", 500, 5.0},
		{3, "Straub AMS", "Описание товара 3", "straub.jpg", "test desc3", 700, 4.5},
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"AMS": AMS,
		})
	})

	r.GET("/product/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "info.html", gin.H{
			"AMS": AMS,
		})

	})

	r.Run(":8080")
}
