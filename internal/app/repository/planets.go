package repository

import "github.com/drakenchef/RIP/internal/app/ds"

func (r *Repository) PlanetsList() (*[]ds.Planet, error) {
	var planets []ds.Planet
	r.db.Preload("is_delete").Where("is_delete = ?", "false").Find(&planets)
	return &planets, nil
}

//func (r *Repository) DeletePlanet(id uint) error {
//	var newStatus ds.Planet
//	r.db.Where("status_name = ?", "уничтожен").First(&newStatus)
//
//	if newStatus.ID != 0 {
//		var planet ds.Planet
//		r.db.First(&planet, id)
//		planet.StatusID = newStatus.ID
//		r.db.Save(&city)
//	}
//	return nil
//}
