package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) UsersList(ctx *gin.Context) {
	users, err := h.Repository.UsersList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "users", users)
}
