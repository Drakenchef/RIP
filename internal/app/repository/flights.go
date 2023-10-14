package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/drakenchef/RIP/internal/app/utils"
)

func (r *Repository) FlightsList() (*[]ds.FlightRequest, error) {
	var flights []ds.FlightRequest
	result := r.db.Preload("User").Where("status !=?", utils.DeletedString).Find(&flights)
	return &flights, result.Error
}
func (r *Repository) UsersFlight() (*[]ds.FlightRequest, error) {
	var flight []ds.FlightRequest
	result := r.db.Preload("User").Preload("PlanetsRequest.Planet").Where("user_id = ?", 1).Find(&flight)
	return &flight, result.Error
}

func (r *Repository) FlightsListByUser(id uint) (*[]ds.FlightRequest, error) {
	var flights []ds.FlightRequest
	result := r.db.Preload("User").Where("user_id = ?", id).Find(&flights)
	return &flights, result.Error
}
func (r *Repository) FlightsListByDate(date string) (*[]ds.FlightRequest, error) {
	var flight []ds.FlightRequest
	result := r.db.Preload("User").Where("date_formation > ?", date).Find(&flight)
	return &flight, result.Error
}

func (r *Repository) DeleteFlight(id uint) error {
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
