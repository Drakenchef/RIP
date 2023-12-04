package repository

import (
	"github.com/drakenchef/RIP/internal/app/ds"
)

func (r *Repository) UsersList() (*[]ds.Users, error) {
	var users []ds.Users
	result := r.db.Find(&users)
	return &users, result.Error
}

func (r *Repository) Register(user *ds.Users) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.Users, error) {
	user := &ds.Users{}

	if err := r.db.Where("login = ?", login).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserById(id uint) *ds.Users {
	user := &ds.Users{}

	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil
	}

	return user
}
