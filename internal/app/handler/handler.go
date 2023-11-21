package handler

import (
	"github.com/drakenchef/RIP/internal/app/config"
	_ "github.com/drakenchef/RIP/internal/app/docs"
	redis2 "github.com/drakenchef/RIP/internal/app/redis"
	"github.com/drakenchef/RIP/internal/app/repository"
	"github.com/drakenchef/RIP/internal/app/role"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.UserCRUD(router)
	h.PlanetCRUD(router)
	h.FlightCRUD(router)
	h.PlanetsRequestsCRUD(router)
	registerStatic(router)
}
func (h *Handler) UserCRUD(router *gin.Engine) {
	router.GET("/users", h.UsersList)
	router.POST("/login", h.Login)
	router.POST("/signup", h.Register)
	router.GET("/logout", h.Logout)
}
func (h *Handler) PlanetCRUD(router *gin.Engine) {
	router.GET("/Planets", h.WithoutAuthCheck(role.Buyer, role.Moder), h.PlanetsList)
	router.GET("/Planet/:id", h.PlanetById)
	router.POST("/Planets", h.WithAuthCheck(role.Moder, role.Admin), h.AddPlanet)
	router.PUT("/Planets/:id", h.WithAuthCheck(role.Moder, role.Admin), h.UpdatePlanet)
	router.DELETE("/Planets", h.WithAuthCheck(role.Moder, role.Admin), h.DeletePlanet)
}
func (h *Handler) FlightCRUD(router *gin.Engine) {
	router.GET("/Flights", h.WithAuthCheck(role.Moder), h.FlightsList)
	router.GET("/Flights/:id", h.WithAuthCheck(role.Moder), h.FlightById)
	router.DELETE("/Flights", h.WithAuthCheck(role.Moder), h.DeleteFlight)
	router.PUT("/Flights", h.WithIdCheck(role.Moder), h.UpdateFlight)
	router.PUT("/FlightsUser/:id", h.WithAuthCheck(role.Buyer, role.Moder), h.UserUpdateFlightStatusById)
	router.PUT("/FlightsModer/:id", h.WithAuthCheck(role.Moder), h.ModerUpdateFlightStatusById)
	router.GET("/UsersFlight", h.WithIdCheck(role.Buyer, role.Moder), h.UsersFlight)
	router.PUT("/UsersFlightUpdate", h.WithIdCheck(role.Buyer, role.Moder), h.UsersUpdateFlight)
}
func (h *Handler) PlanetsRequestsCRUD(router *gin.Engine) {
	router.POST("/PlanetsRequests", h.WithIdCheck(role.Buyer, role.Moder), h.AddPlanetToRequest)
	router.DELETE("/PlanetsRequests", h.WithAuthCheck(role.Buyer, role.Moder), h.DeletePlanetRequest)
	router.PUT("/PlanetsRequests", h.WithAuthCheck(role.Buyer, role.Moder), h.UpdatePlanetNumberInRequest)
	router.GET("/ping", h.WithAuthCheck(role.Moder), h.Ping)
}

//func registerStatic(router *gin.Engine) {
//	router.LoadHTMLGlob("static/html/*")
//	router.Static("/static", "./static")
//	router.Static("/css", "./static")
//	router.Static("/img", "./static")
//}

// request status

func registerStatic(router *gin.Engine) {
	router.GET("/myswagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/static", "./static")
	router.Static("/img", "./static")
}

// MARK: - Error handler
type errorResp struct {
	Status      string `json:"status" example:"error"`
	Description string `json:"description" example:"Описание ошибки"`
}

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

// Ping godoc
// @Summary      Show hello text
// @Description  very friendly response
// @Tags         Tests
// @Security ApiKeyAuth
// @Produce      json
// @Router       /ping [get]
func (h *Handler) Ping(gCtx *gin.Context) {
	name := gCtx.Request.FormValue("name")
	gCtx.String(http.StatusOK, "Hello, %s", name)
}
