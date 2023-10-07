package repository

import "github.com/drakenchef/RIP/internal/app/ds"

func (r *Repository) UsersList() (*[]ds.Users, error) {
	var users []ds.Users
	result := r.db.Find(&users)
	return &users, result.Error
}
