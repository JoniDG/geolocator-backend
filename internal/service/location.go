package service

import (
	"geolocator-backend/internal/repository"
)

type LocationService interface {
	GetLocationByCoords()
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{
		repo: repo,
	}
}

func (r *locationService) GetLocationByCoords() {

}
