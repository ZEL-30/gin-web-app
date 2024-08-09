package repository

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/ZEL-30/gin-web-app/domain"
	"github.com/ZEL-30/gin-web-app/entity"
	"github.com/ZEL-30/gin-web-app/representation"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		db,
	}
}

func (s *userRepo) encodeMD5(value string) string {
	m := md5.New()
	_, _ = m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func (m *userRepo) Register(user entity.User) (*entity.User, error) {
	logrus.Debugf("about to save a user %s", user.Name)
	user.Password = m.encodeMD5(user.Password)
	if err := m.db.Create(&user).Error; err != nil {
		return nil, err
	}
	logrus.Debugf("user %s saved", user.Name)
	return &user, nil
}

func (m *userRepo) Get(id string) (*entity.User, error) {
	logrus.Debugf("about to get a user %s", id)
	var data entity.User
	err := m.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &representation.AppError{
				Code:    http.StatusFound,
				Message: fmt.Sprintf("user %s is not found.", id),
			}
		}
		return &entity.User{}, err
	}

	logrus.Debugf("user %s retrieved", id)
	return &data, err
}

func (m *userRepo) GetAll() ([]*entity.User, error) {
	logrus.Debug("about to get all user")
	var users []*entity.User
	err := m.db.Find(&users).Error
	if err != nil {
		return []*entity.User{}, err
	}
	logrus.Debug("all user retrieved")
	return users, nil
}

func (m *userRepo) Update(user entity.User) (*entity.User, error) {
	logrus.Debugf("about to update a user %s", user.Name)
	err := m.db.Select("name", "updated_at").Updates(&user).Error
	logrus.Debugf("user %s updated", user.Name)
	return &user, err
}

func (m *userRepo) Delete(id string) error {
	logrus.Debugf("about to delete a user %s", id)
	tx := m.db.Begin()
	if err := tx.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	logrus.Debugf("user %s deleted", id)
	return tx.Commit().Error
}
