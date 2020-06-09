package util

import (
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Role int

const (
	NormalUser Role = iota + 1
	GroupAdmin
	SystemAdmin
)

type Claims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserRole Role
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(account string, password string, id int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour * 100)

	claims := Claims{
		EncodeMD5(string(account)),
		EncodeMD5(password),
		NormalUser,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "minus4.cn",
			Id:        strconv.Itoa(id),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func ParseBearerHeader(header string) (token string) {
	s := strings.SplitN(header, ":", 2)
	if len(s) != 2 {
		return
	} else {
		authType := s[0]
		if authType != "Bearer" {
			return ""
		}
		token = s[1]
		return
	}
}
