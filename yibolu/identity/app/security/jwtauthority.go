package security

type JWTAuthority struct {
	signingKey []byte

}

func (j JWTAuthority) GenerateJWT(clientId string, userId string) string {
	return "abc"
}


