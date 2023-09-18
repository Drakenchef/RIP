package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
)

func (r *Repository) PlanetsList() (*[]ds.Planet, error) {
	var planets []ds.Planet
	r.db.Find(&planets)
	return &planets, nil
}
