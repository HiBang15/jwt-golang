package token

import (
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