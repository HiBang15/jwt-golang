package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type JwtClaims struct {
	CompanyId string `json:"company_id,omitempty"`
	Username string `json:"username,omitempty"`
	Roles []int `json:"roles,omitempty"`
	jwt.StandardClaims
}

const ip = "192.168.0.107"

func (claims JwtClaims)Valid() error  {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf("Token invalid!")
}

func (claims JwtClaims)VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
