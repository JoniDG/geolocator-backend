package controller

import (
	"geolocator-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type LocationController interface {
	GetLocationByCoords(ctx *gin.Context)
}

type locationController struct {
	svc service.LocationService
}

func NewLocationController(svc service.LocationService) LocationController {
	return &locationController{
		svc: svc,
	}
}

func (c *locationController) GetLocationByCoords(ctx *gin.Context) {
}
