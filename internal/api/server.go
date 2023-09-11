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

	AMS := []backend.AMS{
		{1, "Новые горизонты", "Россия", "Новые горизонты - это зонд, который был запущен в 2006 году с целью изучения плутония и его спутников. В 2015 году он стал первым и единственным космическим аппаратом, который пролетел мимо Плутона, предоставив уникальные фотографии и данные о планете-карлике. Новые горизонты также изучали другие объекты Кайперова пояса и продолжают свою миссию внутри Солнечной системы.", "newhorizon.jpg", "В добрый путь!", []string{"Марс", "Венера", "Юпитер", "Луна", "комета"}, false},
		{2, "Вояджер", "США", "Вояджер - это космический зонд, запущенный НАСА в 1977 году с целью изучения внешних планет Солнечной системы. Вояджер 1 и Вояджер 2 заложили основу для исследования Юпитера, Сатурна, Урана и Нептуна. Они совершили множество научных открытий о планетах, их спутниках и кольцах, а также помогли расширить наше понимание о Вселенной.", "voyager1.jpg", "Good luck", []string{"Марс", "Венера", "Меркурий", "Уран", "Нептун", "Юпитер", "Сатурн", "Луна", "комета", "астеройд"}, false},
		{3, "Чанъэ", "Китай", "Чанъэ - это межпланетный зонд, запущенный Китаем в 2020 году. Его главная цель - исследование Марса и поиск признаков жизни на красной планете. Космический аппарат оснащен научными приборами, способными изучать атмосферу Марса, его магнитное поле и геологический состав. Чанъэ также отправит ровер на поверхность Марса для сбора образцов грунта и анализа их наличия органических молекул. Ожидается, что эта миссия поможет расширить наше знание о Марсе и его потенциальной пригодности для будущих человеческих колоний.", "chane.jpg", "祝你好運", []string{"Марс", "Меркурий", "Луна"}, false},
	}

	r.GET("/", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/index.html")))
		c.HTML(http.StatusOK, "index.html", gin.H{
			"AMS": AMS,
		})
	})

	r.GET("/AMS/:id", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./templates/info.html")))
		// Получаем значение ID из параметров маршрута
		id := c.Param("id")
		var selectedAMS backend.AMS
		// Здесь мы выбираем конкретный AMS по его ID
		for _, ams := range AMS {
			if strconv.Itoa(ams.ID) == id {
				selectedAMS = ams
				break
			}
		}

		// Здесь мы передаем выбранный AMS в шаблон
		c.HTML(http.StatusOK, "info.html", gin.H{
			"AMS": selectedAMS,
		})
	})
	r.GET("/search", func(c *gin.Context) {
		searchQuery := c.Query("search")

		filteredAMS := []backend.AMS{}
		for _, ams := range AMS {
			if strings.Contains(strings.ToLower(ams.Name), strings.ToLower(searchQuery)) {
				filteredAMS = append(filteredAMS, ams)
			}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"AMS": filteredAMS,
		})
	})

	r.Run(":8080")
}
