package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"time"
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
	PlanetId     uint `json:"Planet_id"`
	FlightNumber uint `json:"flight_number"`
}, userid uint) error {
	var FlightRequest ds.FlightRequest
	var user ds.Users
	r.db.Where("user_id = ? AND status = ?", userid, "создан").First(&FlightRequest)
	r.db.Where("id = ?", userid).First(&user)
	if FlightRequest.ID == 0 {
		newRequest := ds.FlightRequest{UserID: user.ID, UserLogin: user.Login, Status: "создан", DateCreate: time.Now(), Result: "Отсутствует"}
		r.db.Create(&newRequest)
		query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2, $3) ON CONFLICT DO NOTHING;"
		err := r.db.Exec(query, newRequest.ID, pr.PlanetId, pr.FlightNumber)
		if err != nil {
			return err.Error
		}
		return nil
	}
	query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2, $3) ON CONFLICT DO NOTHING;"
	err := r.db.Exec(query, FlightRequest.ID, pr.PlanetId, pr.FlightNumber)
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
