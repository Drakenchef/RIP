package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
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

func (r *Repository) AddPlanetToRequest(pr *struct {
	PlanetId uint `json:"planet_id"`
}) error {
	//var flightRequest ds.FlightRequest
	//if err := r.db.First(&flightRequest, pr.FRID).Error; err != nil {
	//	flightRequest = ds.FlightRequest{ID: pr.FRID, Status: utils.ExistsString, UserID: 1}
	//	r.db.Create(&flightRequest)
	//}
	//
	//query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2,$3);"
	//err := r.db.Exec(query, pr.FRID, pr.PlanetID, pr.FlightNumber)
	//if err != nil {
	//	return err.Error
	//}
	//return nil
	var flightRequest ds.FlightRequest
	r.db.Where("user_id = ?", 1).First(&flightRequest)

	if flightRequest.ID == 0 {
		newRequest := ds.FlightRequest{UserID: 1, Status: "создан"}
		r.db.Create(&newRequest)
		flightRequest = newRequest
	}
	query := "INSERT INTO planets_requests (fr_id, planet_id) VALUES ($1,$2) ON CONFLICT DO NOTHING;"
	err := r.db.Exec(query, flightRequest.ID, pr.PlanetId)
	if err != nil {
		return err.Error
	}
	return nil
}

func (r *Repository) DeletePlanetRequest(frid, planetid uint) error {
	var planetsrequest ds.PlanetsRequest
	err := r.db.Where("fr_id = ? AND planet_id = ?", frid, planetid).Delete(&planetsrequest).Error
	if err != nil {
		return err
	}
	return nil
}
