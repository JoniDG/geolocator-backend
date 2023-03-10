package repository

type LocationRepository interface {
	GetLocationByCoords()
}

type locationRepository struct {
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{}
}

func (r *locationRepository) GetLocationByCoords() {

}
