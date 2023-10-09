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
	//if planetRequest.ID != 0 {
	//	h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
	//	return
	//}
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

func (h *Handler) UpdatePlanetNumberInRequest(ctx *gin.Context) {
	var updatedPlanetRequest ds.PlanetsRequest
	if err := ctx.BindJSON(&updatedPlanetRequest); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedPlanetRequest.FRID == 0 || updatedPlanetRequest.PlanetID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdatePlanetNumberInRequest(&updatedPlanetRequest); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "updated_planet", gin.H{
		"fr_id":         updatedPlanetRequest.FRID,
		"planet_id":     updatedPlanetRequest.PlanetID,
		"flight_number": updatedPlanetRequest.FlightNumber,
	})
}
