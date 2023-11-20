package ds

type FlightsListRes struct {
	Status string          `json:"status"`
	Flight []FlightRequest `json:"flight"`
}

type FlightsListRes2 struct {
	Status string          `json:"status"`
	Flight []FlightRequest `json:"flight"`
}

type DeletePlanetRes struct {
	DeletedId int `json:"deleted_id"`
}

//type AddImageRes struct {
//	Status   string `json:"status"`
//	ImageUrl string `json:"image_url"`
//}

type UpdatedFlightRes struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	DateCreate     string `json:"date_create"`
	DateFormation  string `json:"date_formation"`
	DateCompletion string `json:"date_completion"`
	Status         string `json:"status"`
	AMS            string `json:"ams"`
	UserID         uint   `json:"user_id"`
	ModerID        uint   `json:"moder_id"`
	UserLogin      string `json:"user_login"`
}

type DeleteFlightRes struct {
	Status   string `json:"status"`
	FlightId uint   `json:"flight_id"`
}

type DeleteFlightReq struct {
	ID uint `json:"id"`
}

type UpdateFlightReq struct {
	ID          uint   `json:"id"`
	AMS         string `json:"ams"`
	Description string `json:"description"`
}

type UpdateStatusForModeratorReq struct {
	FlightID uint   `json:"flight_id"`
	Status   string `json:"status"`
}

type UpdateStatusForUserReq struct {
	Status string `json:"status" example:"2"`
}

type DeletePlanetReq struct {
	ID string `json:"id"`
}

type UpdatePlanetReq struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"city_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AddPlanetToRequestReq struct {
	PlanetID int `json:"planet_id" binding:"required" example:"1"`
	//SerialNumber int `json:"serial_number" binding:"required" example:"1"`
}

type AddPlanetToRequestResp struct {
	Status string `json:"status"`
	Id     int    `json:"id"`
}

type UpdatePlanetResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type AddPlanetResp struct {
	Status   string `json:"status"`
	PlanetId string `json:"planet_id"`
}

type PlanetsListResp struct {
	Status  string   `json:"status"`
	Planets []Planet `json:"planets"`
	//BasketId string `json:"basket_id"`
}

//type AddPlaIntoHikeRequest struct {
//}
