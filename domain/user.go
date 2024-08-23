package domain

import (
	"github.com/ZEL-30/gin-web-app/entity"
	"github.com/ZEL-30/gin-web-app/representation"
)

// UserService 用户服务接口
type UserService interface {
	Add(user representation.User) (*representation.User, error)
	Get(id string) (*representation.User, error)
	List() ([]*representation.User, error)
	Update(user representation.User) (*representation.User, error)
	Delete(id string) error
}

// UserRepository 用户存储库接口
type UserRepository interface {
	Add(user entity.User) (*entity.User, error)
	Get(id string) (*entity.User, error)
	List() ([]*entity.User, error)
	Update(user entity.User) (*entity.User, error)
	Delete(id string) error
}
