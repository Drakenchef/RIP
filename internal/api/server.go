package api

import (
	backend "github.com/drakenchef/RIP"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func StartServer() {
	r := gin.Default()

	r.Static("/resources", "./resources")
	r.Static("/static", "./static")

	//МОКИ ЗДЕСЬ
	Planets := []backend.Planets{
		{1, "Марс", "Марс — планета земной группы с разреженной атмосферой (давление на поверхности в 160 раз меньше земного). Особенностями поверхностного рельефа Марса можно считать ударные кратеры наподобие лунных, а также вулканы, долины, пустыни и полярные ледниковые шапки наподобие земных.", "mars.jpg", []string{"М-62", "М-71", "М-73"}},
		{2, "Сатурн", "Сатурн – газовый гигант, наполненный в основном водородом и гелием. Его размеры позволяют разместить в себе 760 планет типа Земли, а масса больше земной в 95 раз. У Сатурна самая низкая плотность. Осевой оборот Сатурна 10 с половиной часов.", "saturn.jpg", []string{"Вояджер-1", "Вояджер-2", "Пионер-11"}},
		{3, "Луна", "Луна – пятый по размеру спутник Солнечной системы. Температура поверхности Луны колеблется от −173 °C ночью до +127 °C в подсолнечной точке. Температура пород на глубине 1 метр постоянна и равна −35 °C. Средний радиус Луны составляет 1737,1 километра, то есть примерно 0,273 радиуса Земли.", "moon.jpg", []string{"Аполлон-11", "Аполлон-12", "Аполлон-14"}},
	}

	r.GET("/", func(c *gin.Context) {
		searchQuery := c.Query("search")

		filteredPlanet := []backend.Planets{}
		for _, c := range Planets {
			if strings.Contains(strings.ToLower(c.Name), strings.ToLower(searchQuery)) {
				filteredPlanet = append(filteredPlanet, c)
			}
		}

		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/mainpage.html")))
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Planet": filteredPlanet,
		})
	})

	r.GET("/Planet/:id", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/planet.html")))
		id := c.Param("id")

		var selectedPlanet backend.Planets

		for _, Planet := range Planets {
			if strconv.Itoa(Planet.ID) == id {
				selectedPlanet = Planet
				break
			}
		}
		//c
		c.HTML(http.StatusOK, "planet.html", gin.H{
			"Planet": selectedPlanet,
		})
	})

	r.Run(":8888")
}
