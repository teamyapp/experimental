module github.com/teamyapp/experimental/yibolu/identity

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/joho/godotenv v1.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
)

// require github.com/teamyapp/protocol/golang/

//replace (
//	github.com/teamyapp/protocol/golang/identity latest => ../protocol/golang/identity
//)
