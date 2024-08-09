package service

import (
	"github.com/ZEL-30/gin-web-app/assembler"
	"github.com/ZEL-30/gin-web-app/domain"
	rep "github.com/ZEL-30/gin-web-app/representation"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserInterface {
	return &userService{
		userRepo,
	}
}

func (s *userService) Register(user rep.User) (*rep.User, error) {
	data, err := s.userRepo.Register(*assembler.UserAsm.ToData(user))
	if err != nil {
		return &rep.User{}, err
	}
	return assembler.UserAsm.ToRepresentation(*data), nil
}

func (s *userService) Get(id string) (*rep.User, error) {
	data, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return assembler.UserAsm.ToRepresentation(*data), nil
}

func (s *userService) GetAll() ([]*rep.User, error) {
	books, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	rtnVal := []*rep.User{}
	for _, user := range books {
		rtnVal = append(rtnVal, assembler.UserAsm.ToRepresentation(*user))
	}
	return rtnVal, nil
}

func (s *userService) Update(user rep.User) (*rep.User, error) {
	data, err := s.userRepo.Update(*assembler.UserAsm.ToData(user))
	if err != nil {
		return nil, err
	}

	return assembler.UserAsm.ToRepresentation(*data), nil
}

func (s *userService) Delete(id string) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
