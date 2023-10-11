package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/drakenchef/RIP/internal/app/utils"
)

func (r *Repository) FlightsList() (*[]ds.FlightRequest, error) {
	var flights []ds.FlightRequest
	result := r.db.Preload("User").Where("status =?", utils.ExistsString).Find(&flights)
	return &flights, result.Error
}

func (r *Repository) FlightById(id uint) (*ds.FlightRequest, error) {
	//flight := ds.FlightRequest{}
	//result := r.db.Preload("User").First(&flight, id)
	//result := r.db.Preload("User").Preload("PlanetRequest.Planet").First(&flight, id)
	//result := r.db.Preload("Planets").First(&flight, id)
	//planets := []ds.Planet{}
	//result1 := r.db.Table("planets").Joins("JOIN planets_requests ON planets.id = planets_requests.planet_id").Where("planets_requests.frid = ?", id).Find(&planets)
	flight := ds.FlightRequest{}
	result := r.db.Preload("User").Preload("PlanetsRequest.Planet").First(&flight, id)

	return &flight, result.Error
	//return &flight, result.Error
}

func (r *Repository) DeleteFlight(id uint) error {
	//flight := ds.FlightRequest{}
	//if result := r.db.Where("status = ?", "удалён").First(&newStatus); result.Error != nil {
	//	return result.Error
	//}
	//hike.StatusID = newStatus.ID
	//if result := r.db.First(&flight, id); result.Error != nil {
	//	return result.Error
	//}
	//result := r.db.Save(&flight)
	//return result.Error
	err := r.db.Model(&ds.FlightRequest{}).Where("id = ?", id).Update("status", utils.DeletedString)
	if err != nil {
		return err.Error
	}
	return nil

}

func (r *Repository) UpdateFlight(updatedFlight *ds.FlightRequest) error {
	oldFlight := ds.FlightRequest{}
	if result := r.db.First(&oldFlight, updatedFlight.ID); result.Error != nil {
		return result.Error
	}
	if updatedFlight.DateCreate.String() != utils.EmptyDate {
		oldFlight.DateCreate = updatedFlight.DateCreate
	}
	if updatedFlight.DateFormation.String() != utils.EmptyDate {
		oldFlight.DateFormation = updatedFlight.DateFormation
	}
	if updatedFlight.DateCompletion.String() != utils.EmptyDate {
		oldFlight.DateCompletion = updatedFlight.DateCompletion
	}
	if updatedFlight.DateApprove.String() != utils.EmptyDate {
		oldFlight.DateApprove = updatedFlight.DateApprove
	}
	if updatedFlight.DateRefuse.String() != utils.EmptyDate {
		oldFlight.DateRefuse = updatedFlight.DateRefuse
	}
	if updatedFlight.Status != "" {
		oldFlight.Status = updatedFlight.Status
	}
	if updatedFlight.AMS != "" {
		oldFlight.AMS = updatedFlight.AMS
	}
	if updatedFlight.UserID != utils.EmptyInt {
		oldFlight.UserID = updatedFlight.UserID
	}
	if updatedFlight.ModerID != utils.EmptyInt {
		oldFlight.ModerID = updatedFlight.ModerID
	}
	*updatedFlight = oldFlight
	result := r.db.Save(updatedFlight)
	return result.Error
}

func (r *Repository) UpdateFlightStatus(updatedFlight *ds.FlightRequest) error {
	oldFlight := ds.FlightRequest{}
	if result := r.db.First(&oldFlight, updatedFlight.ID); result.Error != nil {
		return result.Error
	}
	if updatedFlight.Status != "" {
		oldFlight.Status = updatedFlight.Status
	}
	*updatedFlight = oldFlight
	result := r.db.Save(updatedFlight)
	return result.Error
}
