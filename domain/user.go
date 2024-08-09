package domain

import (
	"github.com/ZEL-30/gin-web-app/entity"
	rep "github.com/ZEL-30/gin-web-app/representation"
)

type UserRepository interface {
	Register(user entity.User) (*entity.User, error)
	Get(id string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Update(user entity.User) (*entity.User, error)
	Delete(id string) error
}

type UserInterface interface {
	Register(user rep.User) (*rep.User, error)
	Get(id string) (*rep.User, error)
	GetAll() ([]*rep.User, error)
	Update(user rep.User) (*rep.User, error)
	Delete(id string) error
}
