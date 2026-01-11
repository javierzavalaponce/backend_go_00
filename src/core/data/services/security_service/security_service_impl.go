package security_service_impl

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SecurityServiceImpl struct {
	secret string
}

func NewSecurityServiceImpl(secret string) *SecurityServiceImpl {
	return &SecurityServiceImpl{secret: secret}
}

func (s *SecurityServiceImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *SecurityServiceImpl) ValidatePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *SecurityServiceImpl) GenerateToken(userID string, duration time.Duration) (string, error) {
	expiration := time.Now().Add(duration).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiration,
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *SecurityServiceImpl) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["user_id"].(string); ok {
			return id, nil
		}
		return "", jwt.ErrInvalidKey
	}

	return "", jwt.ErrInvalidKey
}
