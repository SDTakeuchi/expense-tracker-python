package repository

import "app/models/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
