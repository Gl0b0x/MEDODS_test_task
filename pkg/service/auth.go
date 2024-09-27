package service

import (
	"MEDODS/pkg/models"
	"MEDODS/pkg/repository"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	jwtString string
	repo      repository.Authorization
}

func (s *AuthService) GetUserByUUID(uuid string) (*models.User, error) {
	return s.repo.GetUserByUUID(uuid)
}

func (s *AuthService) UpdateUserRefreshToken(user *models.User, base64RefreshToken string) error {
	hashToken, err := bcryptToken(base64RefreshToken)
	if err != nil {
		return err
	}
	return s.repo.UpdateUserRefreshToken(user.Guid, hashToken)
}

func NewAuthService(repo repository.Authorization, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtString: jwtSecret}
}

func (s *AuthService) generateAccessToken(uuid, ip string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": uuid,
		"ip":  ip,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	//accessTokenString, err := accessToken.SignedString([]byte("jwt_secret"))
	return accessToken.SignedString([]byte(s.jwtString))
}

func (s *AuthService) generateRefreshToken(uuid, ip string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": uuid,
		"ip":  ip,
		"exp": time.Now().Add(time.Minute * 24 * 7).Unix(),
	})
	//accessTokenString, err := accessToken.SignedString([]byte("jwt_secret"))
	return refreshToken.SignedString([]byte(s.jwtString))
}

func (s *AuthService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.jwtString), nil
	})
	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	uuid, err := claims.GetSubject()
	return uuid, err
}

func (s *AuthService) GenerateTokens(user *models.User) (string, string, error) {
	accessToken, err := s.generateAccessToken(user.Guid, *user.Ip)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.generateRefreshToken(user.Guid, *user.Ip)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func EncodeBase64(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(token))
}

func DecodeBase64(tokenString string) (string, error) {
	token, err := base64.StdEncoding.DecodeString(tokenString)
	return string(token), err
}

func bcryptToken(tokenString string) (string, error) {
	if len(tokenString) > 72 {
		tokenString = tokenString[len(tokenString)-72:]
	}
	buff, err := bcrypt.GenerateFromPassword([]byte(tokenString), bcrypt.DefaultCost)
	return string(buff), err
}

func (s *AuthService) CompareHashAndPassword(hashedPassword, password string) error {
	if len(password) > 72 {
		password = password[len(password)-72:]
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
