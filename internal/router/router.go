package router

import (
	"geolocator-backend/internal/controller"
	"geolocator-backend/internal/defines"
	"geolocator-backend/internal/repository"
	"geolocator-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init

	// Repositories init
	repo := repository.NewLocationRepository()

	// Services init
	svc := service.NewLocationService(repo)

	// Controllers init
	ctrl := controller.NewLocationController(svc)

	// Endpoints
	r.GET("/location", ctrl.GetLocationByCoords)

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))
}
