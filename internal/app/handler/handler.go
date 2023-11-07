package handler

import (
	"github.com/drakenchef/RIP/internal/app/config"
	redis2 "github.com/drakenchef/RIP/internal/app/redis"
	"github.com/drakenchef/RIP/internal/app/repository"
	"github.com/drakenchef/RIP/internal/app/role"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.Client
	Config     *config.Config
	Redis      *redis2.Client
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.Client, conf *config.Config, red *redis2.Client) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
		Config:     conf,
		Redis:      red,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/Planets", h.PlanetsList)
	router.GET("/Planet/:id", h.PlanetById)
	router.POST("/Planets", h.AddPlanet)
	router.PUT("/Planets/:id", h.UpdatePlanet)
	router.DELETE("/Planets", h.DeletePlanet)
	//router.POST("/PlanetImage", h.AddImage)

	router.GET("/Flights", h.FlightsList)
	router.GET("/Flights/:id", h.FlightById)
	router.DELETE("/Flights", h.DeleteFlight)
	router.PUT("/Flights", h.UpdateFlight)
	router.PUT("/FlightsUser/:id", h.UserUpdateFlightStatusById)
	router.PUT("/FlightsModer/:id", h.ModerUpdateFlightStatusById)
	router.GET("/UsersFlight", h.UsersFlight)
	//router.PUT("/Flights/:id", h.UpdateFlightStatus)

	router.GET("/users", h.UsersList)
	router.POST("/login", h.Login)
	router.POST("/signup", h.Register)
	router.GET("/logout", h.Logout)

	//router.GET("/PlanetsRequests", h.PlanetsRequestsList)
	router.POST("/PlanetsRequests", h.AddPlanetToRequest)
	router.DELETE("/PlanetsRequests", h.DeletePlanetRequest)
	router.PUT("/PlanetsRequests", h.UpdatePlanetNumberInRequest)

	//router.GET("/users", h.UsersList)
	router.Use(h.WithAuthCheck(role.Manager, role.Admin)).GET("/ping", h.Ping)
	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}

//request status

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	h.Logger.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}

func (h *Handler) successHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		key:      data,
	})
}

func (h *Handler) successAddHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		key:      data,
	})
}

func (h *Handler) Ping(gCtx *gin.Context) {
	name := gCtx.Request.FormValue("name")
	gCtx.String(http.StatusOK, "Hello, %s", name)
}
