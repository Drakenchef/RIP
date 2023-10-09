package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/drakenchef/RIP/internal/app/utils"
)

func (r *Repository) PlanetsRequestsList() (*[]ds.PlanetsRequest, error) {
	var planetsRequest []ds.PlanetsRequest
	result := r.db.Select("fr_id", "planet_id", "flight_number").Find(&planetsRequest)
	return &planetsRequest, result.Error
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

func (r *Repository) AddPlanetToRequest(pr *ds.PlanetsRequest) error {
	var flightRequest ds.FlightRequest
	if err := r.db.First(&flightRequest, pr.FRID).Error; err != nil {
		flightRequest = ds.FlightRequest{ID: pr.FRID, Status: utils.ExistsString, UserID: 1}
		r.db.Create(&flightRequest)
	}

	query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2,$3);"
	err := r.db.Exec(query, pr.FRID, pr.PlanetID, pr.FlightNumber)
	if err != nil {
		return err.Error
	}
	return nil
}
