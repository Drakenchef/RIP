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

func (r *Repository) UpdatePlanetNumberInRequest(updatedPlanetRequest *ds.PlanetsRequest) error {
	var oldPlanetRequest ds.PlanetsRequest
	if result := r.db.First(&oldPlanetRequest, updatedPlanetRequest.FRID, updatedPlanetRequest.PlanetID); result.Error != nil {
		return result.Error
	}
	if updatedPlanetRequest.FlightNumber != 0 {
		oldPlanetRequest.FlightNumber = updatedPlanetRequest.FlightNumber
	}

	*updatedPlanetRequest = oldPlanetRequest
	result := r.db.Save(updatedPlanetRequest)
	return result.Error
}
