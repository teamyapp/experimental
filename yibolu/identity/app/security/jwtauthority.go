package security

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTAuthority struct {
	signingKey []byte
}

var issuer = "Identity"
var TTL = time.Hour * 4

func (j JWTAuthority) GenerateJWT(userID string, clientID string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(TTL)

	claims := IdentityClaims{
		UserID: 	userID,
		ClientID:   clientID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.signingKey)
}

func (j JWTAuthority) DecodeJWT(jwtToken string) (string, string, error) {
	identityClaim, err := j.parse(jwtToken)

	if err != nil {
		return "", "", err
	}

	return identityClaim.UserID, identityClaim.ClientID, nil
}

func (j JWTAuthority) ValidateToken(jwtToken string) bool  {
	_, err := j.parse(jwtToken)

	if err != nil {
		return false
	}

	return true
}

func (j JWTAuthority) parse(jwtToken string) (*IdentityClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(jwtToken, &IdentityClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*IdentityClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func NewJWTAuthority(signingKey []byte) JWTAuthority {
	return JWTAuthority {
		signingKey:		signingKey,
	}
}
