package handler

import (
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) PlanetsRequestsList(ctx *gin.Context) {
	planetsRequests, err := h.Repository.PlanetsRequestsList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "planets_requests", planetsRequests)

}

// UpdatePlanetNumberInRequest godoc
// @Summary Обновление порядка посещения планет
// @Description Обновление порядкого номера планеты в заявке
// @Tags Планеты в заявках
// @Accept json
// @Produce json
// @Param request body ds.UpdatePlanetInRequestNumberReq true "Данные для добавления планеты в заявку"
// @Success 200 {object} ds.UpdatePlanetInRequestNumberRes "Updated Planet In Request ID"
// @Failure 400 {object} errorResp "Плохой запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибку"
// @Router /PlanetsRequests [put]
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

// AddPlanetToRequest godoc
// @Summary Добавление планеты в заявку
// @Security ApiKeyAuth
// @Tags Планеты в заявках
// @Description Добавление планеты в заявку. Если заявка не найдена, она будет сформирована
// @Accept json
// @Produce json
// @Param request body ds.AddPlanetToRequestReq true "Данные для добавления планеты в заявку"
// @Success 200 {object} ds.AddPlanetToRequestResp "ID"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /PlanetsRequests [post]
func (h *Handler) AddPlanetToRequest(ctx *gin.Context) {
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
	//var planetRequest ds.PlanetsRequest
	var request struct {
		PlanetId     uint `json:"Planet_id"`
		FlightNumber uint `json:"flight_number"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.PlanetId == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, fridOrPlanetIsEmpty)
		return
	}

	if err := h.Repository.AddPlanetToRequest(&request, userIDUint); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successAddHandler(ctx, "updated_planet_request", request)
}

// DeletePlanetRequest godoc
// @Summary Удаление планеты из заявки
// @Description Удаление планеты из заявки по идентификатору
// @Tags Планеты в заявках
// @Accept json
// @Produce json
// @Param request body ds.DeletePlanetInRequestReq true "Идентификатор планеты в заявке"
// @Success 200 {object} ds.DeletePlanetInRequestRes "Удаленный идентификатор планеты"
// @Failure 400 {object} errorResp "Плохой запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /PlanetsRequests [delete]
func (h *Handler) DeletePlanetRequest(ctx *gin.Context) {
	var request struct {
		FRID     uint `json:"fr_id"`
		PlanetId uint `json:"planet_id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.FRID == 0 || request.PlanetId == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeletePlanetRequest(request.FRID, request.PlanetId); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "fr_id", request.FRID)
}
