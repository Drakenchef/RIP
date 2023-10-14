package handler

import (
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) FlightsList(ctx *gin.Context) {
	flights, err := h.Repository.FlightsList()

	//user id sort in requests
	if userIdString := ctx.Query("Sort"); userIdString == "ID" {
		var request struct {
			ID uint `json:"id"`
		}
		if err = ctx.BindJSON(&request); err != nil {
			h.errorHandler(ctx, http.StatusBadRequest, err)
			return
		}
		if request.ID == 0 {
			h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
			return
		}
		var flightRequest *[]ds.FlightRequest
		if flightRequest, err = h.Repository.FlightsListByUser(request.ID); err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
		h.successHandler(ctx, "Flights by user id", flightRequest)
		return
	}

	//date sort in requests
	if DateString := ctx.Query("Sort"); DateString == "Date" {
		var request struct {
			DateFormation string `json:"date_formation" time_format:"2006-01-02"`
		}
		if err = ctx.BindJSON(&request); err != nil {
			h.errorHandler(ctx, http.StatusBadRequest, err)
			return
		}
		if request.DateFormation == "" {
			h.errorHandler(ctx, http.StatusBadRequest, errors.New("empty date input"))
			return
		}
		var flightRequest *[]ds.FlightRequest
		if flightRequest, err = h.Repository.FlightsListByDate(request.DateFormation); err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
		h.successHandler(ctx, "Flights by date", flightRequest)
		return
	}

	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Flights", flights)
}

func (h *Handler) UsersFlight(ctx *gin.Context) {
	flight, err := h.Repository.UsersFlight()
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Flight", flight)
}

func (h *Handler) DeleteFlight(ctx *gin.Context) {
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
	if err := h.Repository.DeleteFlight(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Flight_id", request.ID)
	//ctx.Redirect(http.StatusOK, "/Flights")
	//h.FlightsList(ctx)
}

func (h *Handler) UpdateFlight(ctx *gin.Context) {
	var updatedFlight ds.FlightRequest
	if err := ctx.BindJSON(&updatedFlight); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedFlight.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateFlight(&updatedFlight); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "updated_flight", gin.H{
		"id":              updatedFlight.ID,
		"date_create":     updatedFlight.DateCreate,
		"date_formation":  updatedFlight.DateFormation,
		"date_completion": updatedFlight.DateCompletion,
		"date_approve":    updatedFlight.DateApprove,
		"date_refuse":     updatedFlight.DateRefuse,
		"status":          updatedFlight.Status,
		"ams":             updatedFlight.AMS,
		"user_id":         updatedFlight.UserID,
		"moder_id":        updatedFlight.ModerID,
	})
}

func (h *Handler) UpdateFlightById(ctx *gin.Context) {
	var updatedFlight ds.FlightRequest
	if err := ctx.BindJSON(&updatedFlight); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedFlight.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateFlightStatus(&updatedFlight); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "updated_flight", gin.H{
		"id":     updatedFlight.ID,
		"status": updatedFlight.Status,
	})
}

//func (h *Handler) UpdateFlightStatus(ctx *gin.Context) {
//	//id := ctx.Param("id")
//	var updatedFlight ds.FlightRequest
//	if err := ctx.BindJSON(&updatedFlight); err != nil {
//		h.errorHandler(ctx, http.StatusBadRequest, err)
//		return
//	}
//	if updatedFlight.ID == 0 {
//		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
//		return
//	}
//	if err := h.Repository.UpdateFlightStatus(&updatedFlight); err != nil {
//		h.errorHandler(ctx, http.StatusInternalServerError, err)
//		return
//	}
//
//	h.successHandler(ctx, "updated_flight", gin.H{
//		"id":     updatedFlight.ID,
//		"status": updatedFlight.Status,
//	})
//}
