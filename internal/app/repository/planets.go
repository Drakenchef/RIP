package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
	"strconv"
	"strings"
)

func (r *Repository) PlanetsList() (*[]ds.Planet, error) {
	var planets []ds.Planet
	r.db.Find(&planets)
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

func (r *Repository) PlanetById(id string) (*ds.Planet, error) {
	var planets ds.Planet
	intId, _ := strconv.Atoi(id)
	r.db.Find(&planets, intId)
	return &planets, nil
}
