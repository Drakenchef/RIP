package handler

import (
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//	func (h *Handler) FlightsList(ctx *gin.Context) {
//		flights, err := h.Repository.FlightsList()
//
//		//user id sort in requests
//		if userIdString := ctx.Query("user"); userIdString != "" {
//			var request struct {
//				ID uint `json:"id"`
//			}
//			if err = ctx.BindJSON(&request); err != nil {
//				h.errorHandler(ctx, http.StatusBadRequest, err)
//				return
//			}
//			if request.ID == 0 {
//				h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
//				return
//			}
//			var flightRequest *[]ds.FlightRequest
//			if flightRequest, err = h.Repository.FlightsListByUser(request.ID); err != nil {
//				h.errorHandler(ctx, http.StatusInternalServerError, err)
//				return
//			}
//			h.successHandler(ctx, "Flights by user id", flightRequest)
//			return
//		}
//
//		//date sort in requests
//		if DateString := ctx.Query("Sort"); DateString == "Date" {
//			var request struct {
//				DateFormationStart string `json:"date_formation_start" time_format:"2006-01-02"`
//				DateFormationEnd   string `json:"date_formation_end" time_format:"2006-01-02"`
//			}
//			if err = ctx.BindJSON(&request); err != nil {
//				h.errorHandler(ctx, http.StatusBadRequest, err)
//				return
//			}
//			if request.DateFormationStart == "" {
//				h.errorHandler(ctx, http.StatusBadRequest, errors.New("empty date input"))
//				return
//			}
//			if request.DateFormationEnd == "" {
//				h.errorHandler(ctx, http.StatusBadRequest, errors.New("empty date input"))
//				return
//			}
//			var flightRequest *[]ds.FlightRequest
//			if flightRequest, err = h.Repository.FlightsListByDate(request.DateFormationStart, request.DateFormationEnd); err != nil {
//				h.errorHandler(ctx, http.StatusInternalServerError, err)
//				return
//			}
//			h.successHandler(ctx, "Flights by date", flightRequest)
//			return
//		}
//		if StatusString := ctx.Query("Sort"); StatusString == "Status" {
//			var request struct {
//				Status string `json:"status"`
//			}
//			if err = ctx.BindJSON(&request); err != nil {
//				h.errorHandler(ctx, http.StatusBadRequest, err)
//				return
//			}
//			if request.Status == "" {
//				h.errorHandler(ctx, http.StatusBadRequest, errors.New("empty status input"))
//				return
//			}
//			var flightRequest *[]ds.FlightRequest
//			if flightRequest, err = h.Repository.FlightsListByStatus(request.Status); err != nil {
//				h.errorHandler(ctx, http.StatusInternalServerError, err)
//				return
//			}
//			h.successHandler(ctx, "Flights by status", flightRequest)
//			return
//		}
//
//		if err != nil {
//			h.errorHandler(ctx, http.StatusNoContent, err)
//			return
//		}
//		h.successHandler(ctx, "Flights", flights)
//	}
func (h *Handler) FlightsList(ctx *gin.Context) {
	userID := ctx.DefaultQuery("user_id", "")
	datestart := ctx.DefaultQuery("date_formation_start", "")
	dateend := ctx.DefaultQuery("date_formation_end", "")
	status := ctx.DefaultQuery("status", "")

	flights, err := h.Repository.FlightsList(userID, datestart, dateend, status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
		return
	}

	ctx.JSON(http.StatusOK, flights)
}

func (h *Handler) UsersFlight(ctx *gin.Context) {
	// Получение значения userid из контекста
	userID, exists := ctx.Get("user_id")
	if !exists {
		// Обработка ситуации, когда userid отсутствует в контексте
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id not found in context"))
		return
	}

	// Приведение типа, если необходимо
	var userIDUint uint
	switch v := userID.(type) {
	case uint:
		userIDUint = v
	case int:
		userIDUint = uint(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errors.New("failed to convert user_id to uint"))
			return
		}
		userIDUint = uint(i)
	default:
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id is not of a supported type"))
		return
	}

	flight, err := h.Repository.UsersFlight(userIDUint)
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
		"status":          updatedFlight.Status,
		"ams":             updatedFlight.AMS,
		"user_id":         updatedFlight.UserID,
		"moder_id":        updatedFlight.ModerID,
	})
}

func (h *Handler) UserUpdateFlightStatusById(ctx *gin.Context) {
	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	result, err := h.Repository.UserUpdateFlightStatusById(idint)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("can not refactor status"))
		return
	}

	h.successHandler(ctx, "updated_status_by_user", gin.H{
		"id":     result.ID,
		"status": result.Status,
	})
}
func (h *Handler) ModerUpdateFlightStatusById(ctx *gin.Context) {
	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	result, err := h.Repository.ModerUpdateFlightStatusById(idint)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("can not refactor status"))
		return
	}

	h.successHandler(ctx, "updated_status_by_moder", gin.H{
		"id":     result.ID,
		"status": result.Status,
	})
}
func (h *Handler) FlightById(ctx *gin.Context) {
	id := ctx.Param("id")
	flight, err := h.Repository.FlightById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Flight", flight)

}

//
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

//if idint, err := strconv.Atoi(id); err != nil {
//	if idint == 0 {
//		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
//		return
//	} else {
//		var updatedFlight ds.FlightRequest
//		updatedFlight.ID = uint(idint)
//		if err := ctx.BindJSON(&updatedFlight); err != nil {
//			h.errorHandler(ctx, http.StatusBadRequest, err)
//			return
//		}
//
//		if err := h.Repository.UpdateFlightStatus(&updatedFlight); err != nil {
//			h.errorHandler(ctx, http.StatusInternalServerError, err)
//			return
//		}
//
//		h.successHandler(ctx, "updated_flight", gin.H{
//			"id":     updatedFlight.ID,
//			"status": updatedFlight.Status,
//		})
//	}
//} else {
//	h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
//}
