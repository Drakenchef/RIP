package repository

import "github.com/drakenchef/RIP/internal/app/ds"

func (r *Repository) PlanetsRequestsList() (*[]ds.PlanetsRequest, error) {
	var planetsRequest []ds.PlanetsRequest
	result := r.db.Find(&planetsRequest)
	return &planetsRequest, result.Error
}

func (r *Repository) AddPlanetToRequest(pr *ds.PlanetsRequest) error {
	result := r.db.Create(pr)
	return result.Error
}
