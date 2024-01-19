package repository

import (
	"errors"
	"github.com/drakenchef/RIP/internal/app/ds"
	"gorm.io/gorm"
	"time"
)

func (r *Repository) PlanetsRequestsList() (*[]ds.PlanetsRequest, error) {
	var planetsRequest []ds.PlanetsRequest
	result := r.db.Select("fr_id", "planet_id", "flight_number").Find(&planetsRequest)
	return &planetsRequest, result.Error
}

//var oldPlanetRequest ds.PlanetsRequest
//if result := r.db.First(&oldPlanetRequest, updatedPlanetRequest.FRID, updatedPlanetRequest.PlanetID); result.Error != nil {
//	return result.Error
//}
//if updatedPlanetRequest.FlightNumber != 0 {
//	oldPlanetRequest.FlightNumber = updatedPlanetRequest.FlightNumber
//}
//
//*updatedPlanetRequest = oldPlanetRequest
//result := r.db.Save(updatedPlanetRequest)
//return result.Error

func (r *Repository) UpdatePlanetNumberInRequest(updatedPlanetRequest *struct {
	PlanetID     uint   `json:"Planet_id"`
	FRID         uint   `json:"fr_id"`
	FlightNumber uint   `json:"flight_number"`
	Command      string `json:"command"`
}) error {
	if updatedPlanetRequest.Command == "" {
		return errors.New("error with updating planet number in request: Command empty")
	}
	//var flight ds.FlightRequest
	//flight = r.db.Preload("PlanetsRequest.Planet").First(&flight, updatedPlanetRequest.FRID)
	//if len(flight.PlanetsRequest) < 0 {
	//	return errors.New("error with updating planet number in request: No Planets in request")
	//}
	//
	//if updatedPlanetRequest.Command == "up"{
	//	if flight.PlanetsRequest[updatedPlanetRequest.PlanetID]
	//}
	return errors.New("error with updating planet number in request: Command empty")

}

func (r *Repository) AddPlanetToRequest(pr *struct {
	PlanetId uint `json:"Planet_id"`
	//FlightNumber uint `json:"flight_number"`
}, userid uint) error {
	var FlightRequest ds.FlightRequest
	var user ds.Users
	r.db.Where("user_id = ? AND status = ?", userid, "создан").First(&FlightRequest)
	r.db.Where("id = ?", userid).First(&user)
	if FlightRequest.ID == 0 {
		newRequest := ds.FlightRequest{UserID: user.ID, UserLogin: user.Login, Status: "создан", DateCreate: time.Now(), Result: "Отсутствует"}
		r.db.Create(&newRequest)
		query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2, $3) ON CONFLICT DO NOTHING;"
		err := r.db.Exec(query, newRequest.ID, pr.PlanetId, 1)
		if err != nil {
			return err.Error
		}
		return nil
	}

	flightNumber := 0
	r.db.Model(&ds.PlanetsRequest{}).Where("fr_id = ?", FlightRequest.ID).Select("MAX(flight_number)").Scan(&flightNumber)
	query := "INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES ($1,$2, $3) ON CONFLICT DO NOTHING;"
	err := r.db.Exec(query, FlightRequest.ID, pr.PlanetId, flightNumber+1)
	if err != nil {
		return err.Error
	}
	return nil
}

//func (r *Repository) DeletePlanetRequest(frid, planetid uint) error {
//	var planetsrequest ds.PlanetsRequest
//	err := r.db.Where("fr_id = ? AND planet_id = ?", frid, planetid).Delete(&planetsrequest).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (r *Repository) DeletePlanetRequest(frid, planetid uint) error {
	var planetsrequest ds.PlanetsRequest
	err := r.db.Where("fr_id = ? AND planet_id = ?", frid, planetid).First(&planetsrequest).Error
	if err != nil {
		return err
	}
	var flightrequest ds.FlightRequest
	err = r.db.Where("id = ?", planetsrequest.FRID).First(&flightrequest).Error
	if err != nil {
		return err
	}
	// Уменьшаем flight_number у всех planets_request, у которых он больше удаленной записи
	err = r.db.Model(&ds.PlanetsRequest{}).Where("fr_id = ? AND flight_number > ?", frid, planetsrequest.FlightNumber).Update("flight_number", gorm.Expr("flight_number - ?", 1)).Error
	if err != nil {
		return err
	}
	var planetsrequest2 ds.PlanetsRequest
	err2 := r.db.Where("fr_id = ? AND planet_id = ?", frid, planetid).Delete(&planetsrequest2).Error
	if err2 != nil {
		return err2
	}

	return nil
}
