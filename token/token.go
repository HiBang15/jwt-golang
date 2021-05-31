package token

import (
	"fmt"
	"github.com/HiBang15/jwt-golang/model"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	JWTPrivateToken = "SecretTokenSecretToken"
	ip = "192.168.0.107"
)

func GenerateToken(claims *model.JwtClaims, expirationTime time.Time) (string, error)  {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWTPrivateToken))
	if err != nil {
		log.Println("Generate toke fail with err " + err.Error())
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString, origin string) (bool, *model.JwtClaims)  {
	claims := &model.JwtClaims{}
	token, _ := getTokenFromString(tokenString, claims)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func GetClaims(tokenString string) model.JwtClaims  {
	claims := &model.JwtClaims{}

	_, err := getTokenFromString(tokenString, claims)
	if err == nil {
		return *claims
	}
	return *claims
}

func getTokenFromString(tokenString string, claims *model.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWTPrivateToken), nil
	} )
}