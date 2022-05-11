package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *jwtService {
	key := os.Getenv("SECRET_KEY")
	// key := "d19u6s0rnuf8v8thcrawyaaedlqj885577a4e809df7b1f5a0630cfcb429381ed0bd221fde9d540702a193804aa27"

	return &jwtService{
		secretKey: key,
		issuer:    "tfweblog",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("token inválido: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *jwtService) GetIDFromToken(t string) (int64, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid Token: %v", t)
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
		// return []byte("d19u6s0rnuf8v8thcrawyaaedlqj885577a4e809df7b1f5a0630cfcb429381ed0bd221fde9d540702a193804aa27"), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := fmt.Sprint(claims["sum"])
		val, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return 0, err
		}

		return val, nil
	}

	return 0, err
}
