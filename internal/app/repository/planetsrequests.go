package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
)

func (r *Repository) PlanetsRequestsList() (*[]ds.PlanetsRequest, error) {
	var planetsRequest []ds.PlanetsRequest
	result := r.db.Select("fr_id", "planet_id", "flight_number").Find(&planetsRequest)
	return &planetsRequest, result.Error
}

func (r *Repository) AddPlanetToRequest(pr *ds.PlanetsRequest) error {
	query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2,$3);"
	err := r.db.Exec(query, pr.FRID, pr.PlanetID, pr.FlightNumber)
	if err != nil {
		return err.Error
	}
	return nil
}
