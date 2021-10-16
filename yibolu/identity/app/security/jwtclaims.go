package security

import "github.com/dgrijalva/jwt-go"

type IdentityClaims struct {
	UserID 		string
	ClientID 	string
	jwt.StandardClaims
}