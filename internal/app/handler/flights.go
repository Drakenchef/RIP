package handler

import (
	"bytes"
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
	"time"
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
// @Param user_login query string false "Сортировка по login пользователя."
// @Success 200 {array} ds.FlightsListRes "Список заявок"
// @Success 200 {array} ds.FlightsListRes2 "Список заявок"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 204 {object} errorResp "Нет данных"
// @Router /Flights [get]
func (h *Handler) FlightsList(ctx *gin.Context) {
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
	var thisUser *ds.Users
	thisUser = h.Repository.GetUserById(userIDUint)
	if thisUser.Role == 0 {
		flight, err := h.Repository.UsersFlight(userIDUint)
		if err != nil {
			h.errorHandler(ctx, http.StatusNoContent, err)
			return
		}
		h.successHandler(ctx, "Flight", flight)

	} else {
		userlogin := ctx.DefaultQuery("user_login", "")
		datestart := ctx.DefaultQuery("date_formation_start", "")
		dateend := ctx.DefaultQuery("date_formation_end", "")
		status := ctx.DefaultQuery("status", "")

		flights, err := h.Repository.FlightsList(userlogin, datestart, dateend, status)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
			return
		}

		type FlightResponse struct {
			ID             uint      `json:"id"`
			DateCreate     time.Time `json:"date_create"`
			DateFormation  time.Time `json:"date_formation"`
			DateCompletion time.Time `json:"date_completion"`
			Status         string    `json:"status"`
			Ams            string    `json:"ams,omitempty"`
			ModerLogin     string    `json:"moder_login"`
			UserLogin      string    `json:"user_login"`
			Result         string    `json:"Result"`
		}

		flightResponses := []FlightResponse{}
		for _, flight := range *flights {
			flightResponse := FlightResponse{
				ID:             flight.ID,
				DateCreate:     flight.DateCreate,
				DateFormation:  flight.DateFormation,
				DateCompletion: flight.DateCompletion,
				Status:         flight.Status,
				ModerLogin:     flight.ModerLogin,
				UserLogin:      flight.UserLogin,
				Result:         flight.Result,
			}
			flightResponses = append(flightResponses, flightResponse)
		}

		// Отправка измененного JSON-ответа без user_id и moder_id
		//ctx.JSON(http.StatusOK, flightResponses)
		h.successHandler(ctx, "Flights", flightResponses)

	}
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
// @Param updatedFLight body ds.UpdateFlightReq true "Данные для обновления заявки"
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
// @Param id path string true "ID заявки"
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
	// Создаем структуру для запроса
	requestBody, err := json.Marshal(map[string]string{
		"flight_id": id,
	})
	if err != nil {
		// Обработка ошибки маршалинга JSON
		ctx.String(http.StatusInternalServerError, "Error creating request body: %v", err)
		return
	}

	// Отправляем запрос на внешний сервис
	resp, err := http.Post("http://127.0.0.1:8000/start-async-update/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		// Обработка ошибки выполнения запроса
		ctx.String(http.StatusInternalServerError, "Error sending request to the external service: %v", err)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		// Обработка случая, когда внешний сервис вернул ошибку
		ctx.String(resp.StatusCode, "External service returned: %s", resp.Status)
		return
	}

	// Все хорошо, возвращаем HTTP статус 200 OK
	ctx.Status(http.StatusOK)
}

// ModerUpdateFlightStatusById godoc
// @Summary Обновление статуса заявки для модератора
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление статуса заявки для модератора.
// @Accept json
// @Produce json
// @Param id path string true "ID заявки"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /FlightsModer/{id} [put]
func (h *Handler) ModerUpdateFlightStatusById(ctx *gin.Context) {
	moderID, exists := ctx.Get("user_id")
	if !exists {
		// Обработка ситуации, когда userid отсутствует в контексте
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("moder_id not found in context"))
		return
	}
	// Приведение типа, если необходимо
	var moderIDUint uint
	switch v := moderID.(type) {
	case uint:
		moderIDUint = v
	case int:
		moderIDUint = uint(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errors.New("failed to convert moder_id to uint"))
			return
		}
		moderIDUint = uint(i)
	default:
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("moder_id is not of a supported type"))
		return
	}

	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	result, err := h.Repository.ModerUpdateFlightStatusById(idint, moderIDUint)
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

func (h *Handler) UpdateFlightAsyncResult(ctx *gin.Context) {
	var req struct {
		Result     string `json:"result"`
		AccessHash string `json:"access_hash"`
	}

	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := ctx.BindJSON(&req); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if req.AccessHash == "ASLDKjalksdjalskjdlk12lk3jfjkfdsfdasdASIODU*As" {
		if err := h.Repository.UpdateFlightAsyncResult(idint, req.Result); err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
		h.successHandler(ctx, "result_updated", gin.H{
			"satellite_id": idint,
			"Result":       req.Result,
		})
	} else {
		h.errorHandler(ctx, http.StatusForbidden, errors.New(req.AccessHash))
		return
	}
}
