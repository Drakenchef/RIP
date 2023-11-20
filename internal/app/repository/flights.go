package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/drakenchef/RIP/internal/app/utils"
)

func (r *Repository) FlightsList(userID, datestart, dateend, status string) (*[]ds.FlightRequest, error) {
	var flights []ds.FlightRequest
	db := r.db.Preload("User").Where("status !=?", utils.DeletedString)

	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}

	if datestart != "" && dateend != "" {
		db = db.Where("date_formation > ? AND date_formation < ?", datestart, dateend)
	}

	if status != "" {
		db = db.Where("status = ?", status)
	}
	for i := range flights {
		flights[i].UserLogin = flights[i].User.Login
	}
	result := db.Find(&flights)
	return &flights, result.Error
}
func (r *Repository) UsersFlight(userid uint) (*[]ds.FlightRequest, error) {
	var flight []ds.FlightRequest
	result := r.db.Preload("User").Preload("PlanetsRequest.Planet").Where("user_id = ?", userid).Find(&flight)
	return &flight, result.Error
}

func (r *Repository) FlightsListByStatus(status string) (*[]ds.FlightRequest, error) {
	var flights []ds.FlightRequest
	result := r.db.Preload("User").Where("status = ?", status).Find(&flights)
	return &flights, result.Error
}

func (r *Repository) DeleteFlight(id uint) error {
	err := r.db.Model(&ds.FlightRequest{}).Where("id = ?", id).Update("status", utils.DeletedString)
	if err != nil {
		return err.Error
	}
	return nil

}
func (r *Repository) UsersUpdateFlight(updatedFlight *ds.FlightRequest, userid uint) error {
	oldFlight := ds.FlightRequest{}
	result := r.db.Where("user_id = ?", userid).Find(&oldFlight)
	if result.Error != nil {
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
	if updatedFlight.Status != "" {
		if updatedFlight.Status == "в работе" && oldFlight.Status == "создан" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "завёршён" && oldFlight.Status == "в работе" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "удалён" && oldFlight.Status == "отменён" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "отменён" && oldFlight.Status != "удалён" {
			oldFlight.Status = updatedFlight.Status
		}
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
	result = r.db.Save(updatedFlight)
	return result.Error
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
	if updatedFlight.Status != "" {
		if updatedFlight.Status == "в работе" && oldFlight.Status == "создан" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "завёршён" && oldFlight.Status == "в работе" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "удалён" && oldFlight.Status == "отменён" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "отменён" && oldFlight.Status != "удалён" {
			oldFlight.Status = updatedFlight.Status
		}
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
		if updatedFlight.Status == "в работе" && oldFlight.Status == "создан" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "завёршён" && oldFlight.Status == "в работе" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "удалён" && oldFlight.Status == "отменён" {
			oldFlight.Status = updatedFlight.Status
		}
		if updatedFlight.Status == "отменён" && oldFlight.Status != "удалён" {
			oldFlight.Status = updatedFlight.Status
		}

	}
	*updatedFlight = oldFlight
	result := r.db.Save(updatedFlight)
	return result.Error
}

func (r *Repository) UserUpdateFlightStatusById(id int) (*ds.FlightRequest, error) {
	var flight ds.FlightRequest
	result := r.db.First(&flight, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Меняем статус тут
	if flight.Status == "создан" {
		flight.Status = "в работе"
	} else if flight.Status == "в работе" {
		flight.Status = "отменён"
	}

	// Сохраняем изменения в базе данных
	if err := r.db.Save(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}
func (r *Repository) ModerUpdateFlightStatusById(id int) (*ds.FlightRequest, error) {
	var flight ds.FlightRequest
	result := r.db.First(&flight, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Меняем статус тут
	if flight.Status == "отменён" {
		flight.Status = "удалён"
	}

	// Сохраняем изменения в базе данных
	if err := r.db.Save(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func (r *Repository) FlightById(id string) (*ds.FlightRequest, error) {
	//var flight ds.FlightRequest
	//intId, _ := strconv.Atoi(id)
	//r.db.Find(&flight, intId)
	//return &flight, nil
	var flight ds.FlightRequest
	result := r.db.Preload("PlanetsRequest.Planet").Where("id", id).Find(&flight)
	return &flight, result.Error
}
