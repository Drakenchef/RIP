package handler

import (
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FlightsList godoc
// @Summary Список заявок
// @Tags Заявки
// @Security ApiKeyAuth
// @Description Получение списка заявок с фильтрами по статусу, дате начала и дате окончания, пользователю.
// @Produce json
// @Param status query string false "Статус заявки."
// @Param date_formation_start query string false "Дата начала периода фильтрации в формате '2006-01-02'."
// @Param date_formation_end query string false "Дата окончания периода фильтрации в формате '2006-01-02'."
// @Success 200 {array} ds.FlightsListRes "Список заявок"
// @Success 200 {array} ds.FlightsListRes2 "Список заявок"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 204 {object} errorResp "Нет данных"
// @Router /Flights [get]
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

// UsersFlight godoc
// @Summary Список заявок пользователя
// @Tags Заявки
// @Security ApiKeyAuth
// @Description Получение списка заявок пользователем.
// @Produce json
// @Success 200 {array} ds.FlightsListRes "Список заявок"
// @Success 200 {array} ds.FlightsListRes2 "Список заявок"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 204 {object} errorResp "Нет данных"
// @Router /UsersFlight [get]
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

// DeleteFlight godoc
// @Summary Удаление заявки
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Удаление заявки по идентификатору.
// @Accept json
// @Produce json
// @Param request body ds.DeleteFlightReq true "Идентификатор заявки для удаления"
// @Success 200 {object} ds.DeleteFlightRes "Успешное удаление заявки"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Flights [delete]
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

// UpdateFlight godoc
// @Summary Обновление данных о заявке
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление данных о заявке.
// @Accept json
// @Produce json
// @Param updatedHike body ds.UpdateFlightReq true "Данные для обновления заявки"
// @Success 200 {object} ds.UpdatedFlightRes "Успешное обновление данных о заявке"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Flights [put]
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

// UsersUpdateFlight godoc
// @Summary Обновление данных о заявке пользователем
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление данных о заявке пользователем.
// @Accept json
// @Produce json
// @Param updatedHike body ds.UpdateFlightReq true "Данные для обновления заявки пользователем"
// @Success 200 {object} ds.UpdatedFlightRes "Успешное обновление данных о заявке пользователя"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /UsersFlightUpdate [put]
func (h *Handler) UsersUpdateFlight(ctx *gin.Context) {
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

	var updatedFlight ds.FlightRequest
	if err := ctx.BindJSON(&updatedFlight); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.Repository.UsersUpdateFlight(&updatedFlight, userIDUint); err != nil {
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

// UserUpdateFlightStatusById godoc
// @Summary Обновление статуса заявки для пользователя.
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление статуса заявки для пользователя.
// @Accept json
// @Produce json
// @Param body body ds.UpdateStatusForUserReq true "Детали обновления статуса"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /FlightsUser/{id} [put]
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

// ModerUpdateFlightStatusById godoc
// @Summary Обновление статуса заявки для модератора
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление статуса заявки для модератора.
// @Accept json
// @Produce json
// @Param body body ds.UpdateStatusForModeratorReq true "Детали обновления статуса"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /FlightsModer/{id} [put]
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

// FlightById godoc
// @Summary Получение информации о заявке по её ID.
// @Tags Заявки
// @Description Получение информации о заявке по его ID.
// @Produce json
// @Param id path string true "ID заявки"
// @Success 200 {object} ds.FlightsListRes2 "Информация о заявке по ID"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 404 {object} errorResp "Заявка не найдена"
// @Router /Flights/{id} [get]
func (h *Handler) FlightById(ctx *gin.Context) {
	id := ctx.Param("id")
	flight, err := h.Repository.FlightById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Flight", flight)

}
