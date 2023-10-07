package handler

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PlanetsRequestsList(ctx *gin.Context) {
	planetsRequests, err := h.Repository.PlanetsRequestsList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "planets_requests", planetsRequests)
}

func (h *Handler) AddPlanetToRequest(ctx *gin.Context) {
	var planetRequest ds.PlanetsRequest
	if err := ctx.BindJSON(&planetRequest); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if planetRequest.ID != 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
		return
	}
	if planetRequest.FlightNumber == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, flightNumberCannotBeEmpty)
		return
	}
	if planetRequest.FRID == 0 || planetRequest.PlanetID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, fridOrPlanetIsEmpty)
		return
	}

	if err := h.Repository.AddPlanetToRequest(&planetRequest); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successAddHandler(ctx, "updated_planet_request", planetRequest)
}
