package auth

import (
	"crypto/subtle"
	"net/http"
	"strings"
	"time"

	"github.com/ZEL-30/gin-web-app/domain"
	rep "github.com/ZEL-30/gin-web-app/representation"
	"github.com/ZEL-30/gin-web-app/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthService(db *gorm.DB) *authService {
	return &authService{
		db:        db,
		jwtSecret: "gin-zel",
	}
}

func (s *authService) Register(user rep.User) error {
	user.Password = util.EncodeMD5(user.Password)
	if err := s.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *authService) Auth(username string, password string) error {
	var user struct {
		Password string
	}
	// Fetch user from database
	if err := s.db.Table("user").Where("name = ?", username).Take(&user).Error; err != nil {
		return &rep.AppError{
			Code:    http.StatusUnauthorized,
			Message: "User does not exist",
		}
	}

	// Verify password
	if !s.verifyPassword(password, user.Password) {
		return &rep.AppError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid username or password",
		}
	}

	return nil
}

func (s *authService) GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := domain.Claims{
		Username: username,
		Password: util.EncodeMD5(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    s.jwtSecret,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) ParseToken(token string) (*domain.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*domain.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (s *authService) GetUserFromToken(token string) (string, error) {
	claims, err := s.ParseToken(token)
	if err != nil {
		return "", err
	}

	return claims.Username, nil
}

func (s *authService) ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s *authService) verifyPassword(password, hash string) bool {
	return subtle.ConstantTimeCompare([]byte(util.EncodeMD5(password)), []byte(hash)) == 1
}
