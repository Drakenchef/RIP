package handler

import (
	"github.com/drakenchef/RIP/internal/app/ds"
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
		ctx.JSON(http.StatusOK, gin.H{
			"Planets": planets,
		})
	} else {

		filteredPlanets, err := h.Repository.SearchPlanet(searchQuery)
		if err != nil {
			// обработка ошибки
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Planets": filteredPlanets,
		})

	}
}

func (h *Handler) planetById(ctx *gin.Context) {
	id := ctx.Param("id")
	planets, err := h.Repository.PlanetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Planets": planets,
	})
}

func (h *Handler) DeletePlanet(ctx *gin.Context) {
	var request struct {
		ID uint `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeletePlanet(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Planet_id", request.ID)
}

func (h *Handler) AddPlanet(ctx *gin.Context) {
	var newPlanet ds.Planet
	if err := ctx.BindJSON(&newPlanet); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if newPlanet.ID != 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
		return
	}
	if newPlanet.Name == "" {
		h.errorHandler(ctx, http.StatusBadRequest, planetCannotBeEmpty)
		return
	}
	if err := h.Repository.AddPlanet(&newPlanet); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/Planets")
}

func (h *Handler) UpdatePlanet(ctx *gin.Context) {
	var updatedPlanet ds.Planet
	if err := ctx.BindJSON(&updatedPlanet); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedPlanet.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdatePlanet(&updatedPlanet); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "updated_planet", gin.H{
		"id":          updatedPlanet.ID,
		"name":        updatedPlanet.Name,
		"description": updatedPlanet.Description,
		"radius":      updatedPlanet.Radius,
		"distance":    updatedPlanet.Distance,
		"gravity":     updatedPlanet.Gravity,
		"image":       updatedPlanet.Image,
		"type":        updatedPlanet.Type,
		"is_delete":   updatedPlanet.IsDelete,
	})
}
