package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/romankravchuk/toronto-pizza/internal/repository"
	dto "github.com/romankravchuk/toronto-pizza/internal/router/handlers/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	rep repository.IUserRepository
}

func NewAuthService(rep repository.IUserRepository) IAuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) GetUserByPhone(ctx context.Context, phone string) (*dto.UserDTO, error) {
	user, err := s.rep.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	userDto := dto.NewUserDTO(user.ID, user.Name, user.Phone)
	return userDto, nil
}

func (s *AuthService) ValidatePassword(phone string, pwd string) error {
	user, err := s.rep.GetByPhone(context.Background(), phone)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(pwd))
}

func (s *AuthService) CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (s *AuthService) ValidateToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
