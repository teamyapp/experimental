package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	JWTSigningKey      	string `envconfig:"JWT_SIGNING_KEY"`
	CaesarCipherOffset 	int    `envconfig:"CAESAR_CIPHER_OFFSET"`
	OAuthProviders 		int    `envconfig:"CAESAR_CIPHER_OFFSET"`
	GoogleClientID     	string `envconfig:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret 	string `envconfig:"GOOGLE_CLIENT_SECRET"`
	JwtTTL				int    `envconfig:"JWT_TTL"`
	//DbHost             string `envconfig:"DB_HOST"`
	//DbPort             int    `envconfig:"DB_PORT"`
	//DBName             string `envconfig:"DB_NAME"`
	//DbUser             string `envconfig:"DB_USER"`
	//DbPassword         string `envconfig:"DB_PASSWORD"`
	WebAPIPort         	int    `envconfig:"WEB_API_PORT"`
	//GRPCAPIPort        int    `envconfig:"GRPC_API_PORT"`
}

func FromEnv() Config {
	err := autoLoadEnv()
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
	return config
}

func autoLoadEnv() error {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return nil
	}

	return godotenv.Load()
}
