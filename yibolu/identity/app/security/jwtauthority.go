package security

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTAuthority struct {
	signingKey 	[]byte
	issuer		string
	ttl			time.Duration
}

func (j JWTAuthority) GenerateJWT(payload interface{}) (string, error) {
	payloadMap, err := toMap(payload)
	if err != nil {
		return "", err
	}

	j.loadStandardInfo(payloadMap)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payloadMap))
	return token.SignedString(j.signingKey)
}

func (j JWTAuthority) DecodeJWT(jwtToken string, output interface{})  error {
	claims, err := j.parse(jwtToken)

	if err != nil {
		return err
	}

	buf, err := json.Marshal(map[string]interface{}(claims))
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, output)
}

func (j JWTAuthority) ValidateToken(jwtToken string) bool  {
	claims, err := j.parse(jwtToken)

	if err != nil {
		return false
	}

	return checkTokenRemainingValidity(claims["exp"])
}

func checkTokenRemainingValidity(timestamp interface{}) bool {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remain := tm.Sub(time.Now())
		if remain > 0 {
			return true
		}
	}
	return false
}

func (j JWTAuthority) parse(jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("token payload is not map")
	}

	return claims, nil
}

func (j JWTAuthority) loadStandardInfo(payloadMap map[string]interface{})  {
	payloadMap["iat"] = time.Now()
	payloadMap["issuer"] = j.issuer
	payloadMap["exp"] = time.After(j.ttl)
}

func toMap(input interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	jsonBuf, _ := json.Marshal(input)
	err := json.Unmarshal(jsonBuf, &output)
	return output, err
}


func NewJWTAuthority(signingKey []byte, issuer string, ttl time.Duration) JWTAuthority {
	return JWTAuthority {
		signingKey:		signingKey,
		issuer:			issuer,
		ttl: 			ttl,
	}
}
