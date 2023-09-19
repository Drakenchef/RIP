package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PlanetsList(ctx *gin.Context) {
	searchQuery := ctx.Query("search")
	if searchQuery == "" {
		planets, err := h.Repository.PlanetsList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Planets": planets,
		})
	} else {

		filteredPlanets, err := h.Repository.SearchPlanet(searchQuery)
		if err != nil {
			// обработка ошибки
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Planets": filteredPlanets,
		})

	}
}

func (h *Handler) PlanetById(ctx *gin.Context) {
	id := ctx.Param("id")
	planets, err := h.Repository.PlanetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "info.html", gin.H{
		"Planets": planets,
	})
}

func (h *Handler) DeletePlanet(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Repository.DeletePlanet(id)
	ctx.Redirect(http.StatusFound, "/")
}
