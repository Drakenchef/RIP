package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"strings"
)

func (r *Repository) PlanetsList() (*[]ds.Planet, error) {
	var planets []ds.Planet
	r.db.Where("is_delete = ?", false).Find(&planets)
	return &planets, nil
}

func (r *Repository) SearchPlanet(search string) (*[]ds.Planet, error) {
	var planets []ds.Planet
	r.db.Find(&planets)

	var filteredPlanets []ds.Planet
	for _, planet := range planets {
		if strings.Contains(strings.ToLower(planet.Name), strings.ToLower(search)) {
			filteredPlanets = append(filteredPlanets, planet)
		}
	}

	return &filteredPlanets, nil
}

func (r *Repository) PlanetById(id uint) (*ds.Planet, error) {
	var planets ds.Planet
	r.db.Find(&planets, id)
	return &planets, nil
}

// asd
func (r *Repository) DeletePlanet(id uint) error {
	//query := "UPDATE Planets SET is_delete = true WHERE id = $1"
	//r.db.Exec(query, id)
	err := r.db.Model(&ds.Planet{}).Where("id = ?", id).Update("is_delete", true)
	if err != nil {
		return err.Error
	}
	return nil
}
func (r *Repository) AddPlanet(planet *ds.Planet) error {
	result := r.db.Create(&planet)
	return result.Error
}
func (r *Repository) UpdatePlanet(updatedPlanet *ds.Planet) error {
	var oldPlanet ds.Planet
	if result := r.db.First(&oldPlanet, updatedPlanet.ID); result.Error != nil {
		return result.Error
	}
	if updatedPlanet.Name != "" {
		oldPlanet.Name = updatedPlanet.Name
	}
	if updatedPlanet.Description != "" {
		oldPlanet.Description = updatedPlanet.Description
	}
	if updatedPlanet.Radius != 0 {
		oldPlanet.Radius = updatedPlanet.Radius
	}
	if updatedPlanet.Distance != 0 {
		oldPlanet.Distance = updatedPlanet.Distance
	}
	if updatedPlanet.Gravity != 0 {
		oldPlanet.Gravity = updatedPlanet.Gravity
	}
	//if updatedPlanet.Image != "" {
	//	oldPlanet.Image = updatedPlanet.Image
	//}
	if updatedPlanet.Type != "" {
		oldPlanet.Type = updatedPlanet.Type
	}
	if updatedPlanet.IsDelete != true {
		oldPlanet.IsDelete = updatedPlanet.IsDelete
	}
	*updatedPlanet = oldPlanet
	result := r.db.Save(updatedPlanet)
	return result.Error
}

func (r *Repository) UpdatePlanetImage(id string, newImageURL string) error {
	planet := ds.Planet{}
	if result := r.db.First(&planet, id); result.Error != nil {
		return result.Error
	}
	planet.Image = newImageURL
	result := r.db.Save(planet)
	return result.Error
}
