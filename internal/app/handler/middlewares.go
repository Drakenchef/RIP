package handler

import (
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/drakenchef/RIP/internal/app/role"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

const jwtPrefix = "Bearer "

func (h *Handler) WithAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		jwtStr := gCtx.GetHeader("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}

		jwtStr = jwtStr[len(jwtPrefix):]
		err := h.Redis.CheckJWTInBlacklist(gCtx.Request.Context(), jwtStr)
		if err == nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !errors.Is(err, redis.Nil) {
			gCtx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.Config.JWT.Token), nil
		})
		if err != nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			h.Logger.Error(err)
			return
		}

		myClaims := token.Claims.(*ds.JWTClaims)

		for _, oneOfAssignedRole := range assignedRoles {
			if myClaims.Role == oneOfAssignedRole {
				//gCtx.Next()
				return
				//.
			}
		}
		gCtx.AbortWithStatus(http.StatusForbidden)
		h.Logger.Infof("role %s is not assigned in %s", myClaims.Role, assignedRoles)
		return
	}
}

func (h *Handler) WithIdCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		jwtStr := gCtx.GetHeader("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}

		jwtStr = jwtStr[len(jwtPrefix):]
		err := h.Redis.CheckJWTInBlacklist(gCtx.Request.Context(), jwtStr)
		if err == nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !errors.Is(err, redis.Nil) {
			gCtx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.Config.JWT.Token), nil
		})
		if err != nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			h.Logger.Error(err)
			return
		}

		myClaims := token.Claims.(*ds.JWTClaims)
		for _, oneOfAssignedRole := range assignedRoles {
			if myClaims.Role == oneOfAssignedRole {
				gCtx.Set("user_id", myClaims.UserID)
				gCtx.Next()
				return
			}
		}
		gCtx.AbortWithStatus(http.StatusForbidden)
		h.Logger.Infof("role %s is not assigned in %s", myClaims.Role, assignedRoles)
		return
	}

}
func (h *Handler) WithoutAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		jwtStr := gCtx.GetHeader("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			gCtx.Set("user_id", 0)
			return
		}

		jwtStr = jwtStr[len(jwtPrefix):]
		err := h.Redis.CheckJWTInBlacklist(gCtx.Request.Context(), jwtStr)
		if err == nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !errors.Is(err, redis.Nil) {
			gCtx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.Config.JWT.Token), nil
		})
		if err != nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			h.Logger.Error(err)
			return
		}

		myClaims := token.Claims.(*ds.JWTClaims)
		for _, oneOfAssignedRole := range assignedRoles {
			if myClaims.Role == oneOfAssignedRole {
				gCtx.Set("user_id", myClaims.UserID)
				gCtx.Next()
				return
			}
		}
		gCtx.AbortWithStatus(http.StatusForbidden)
		h.Logger.Infof("role %s is not assigned in %s", myClaims.Role, assignedRoles)
		return
	}

}
