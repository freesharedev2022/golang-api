package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_JWT"))

type Claims struct {
	Name  string `json: name`
	Email string `json: email`
	Id    uint   `json: id`
	jwt.StandardClaims
}

func GenerateJWT(email string, name string, id uint) (tokenString string, exp uint) {
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &Claims{
		Email: email,
		Name:  name,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ = token.SignedString(jwtKey)
	exp = uint(expirationTime.Unix())
	return
}

func ValidateToken(signedToken string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, err
}
