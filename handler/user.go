package handler

import (
	"net/http"

	"github.com/ZEL-30/gin-web-app/domain"
	rep "github.com/ZEL-30/gin-web-app/representation"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type userHandler struct {
	userService domain.UserService
	authService domain.AuthInterface
}

func NewUserHandler(userService domain.UserService, authService domain.AuthInterface) userHandler {
	return userHandler{
		userService,
		authService,
	}
}

func (s *userHandler) Add(c *gin.Context) {
	user := &rep.User{}
	if err := c.ShouldBind(user); err != nil {
		appErr := &rep.AppError{
			Code:    http.StatusBadRequest,
			Message: Message.InvalidJson,
		}
		_ = c.Error(appErr)
		return
	}
	logger.Debugf("Received request to add a user %s.", user.Name)

	token := s.authService.ExtractToken(c)
	user.Name, _ = s.authService.GetUserFromToken(token)

	rtnVal, err := s.userService.Add(*user)
	if err != nil {
		_ = c.Error(err)
		return
	}
	logger.Debugf("The user %s is added successfully.", rtnVal.Name)
	c.JSON(http.StatusCreated, rtnVal)
}

func (s *userHandler) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := s.userService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *userHandler) List(c *gin.Context) {
	users, err := s.userService.List()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *userHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_, err := s.userService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	user := &rep.User{}
	if err := c.ShouldBind(user); err != nil {
		_ = c.Error(&rep.AppError{
			Code:    http.StatusBadRequest,
			Message: Message.InvalidJson,
		})
		return
	}
	logger.Debugf("Received request to add a user %s", user.Name)

	user, err = s.userService.Update(*user)
	if err != nil {
		_ = c.Error(err)
		return
	}
	logger.Debugf("The user %s is updated successfully.", user.Name)
	c.JSON(http.StatusOK, user)
}

func (s *userHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	logger.Debugf("Received request to delete a user %s.", id)
	err := s.userService.Delete(id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	logger.Debugf("The user %s is deleted successfully.", id)
	c.JSON(http.StatusNoContent, nil)
}
