package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/models"
)

func (u *UserService) GenerateToken(userEmail string) (string, error) {
	user, err := u.repo.FindUserByEmail(userEmail)
	if err != nil {
		return "", err
	}

	claims := &models.Claims{
		UserID: int(user.UserID),
		Email:  user.Email,
		Role:   user.Role,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("q3e67yajhsdb4")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}