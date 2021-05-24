package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

//Environment -- Env
type Environment struct {
	ApiPort             int    `env:"API_PORT"`
	ApiSecret           string `env:"API_SECRET"`
	AwsEndpoint         string `env:"AWS_ENDPOINT"`
	AwsRegion           string `env:"AWS_REGION"`
	DBDriver            string `env:"DB_DRIVER"`
	DBUser              string `env:"DB_USER"`
	DBPassword          string `env:"DB_PASSWORD"`
	DBPort              string `env:"DB_PORT"`
	DBHost              string `env:"DB_HOST"`
	DBName              string `env:"DB_NAME"`
	SQSQueueURL         string `env:"SQS_QUEUE_URL"`
	SQSQueueSleepTime   string `env:"SQS_QUEUE_SLEEP_TIME"`
	TokenExpirationTime int    `env:"TOKEN_EXPIRATION_TIME"`
}

//ENV -- OUT
var ENV Environment

//Start -- Get enviroment variables
func Start() {
	_, err := env.UnmarshalFromEnviron(&ENV)
	if err != nil {
		log.Fatal(err)
	}
}
