package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PlanetsList(ctx *gin.Context) {
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
}
