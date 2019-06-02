package env

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Env struct {
	AccessToken string `envconfig:"LINE_ACCESS_TOKEN" required:"true"`
	ChannelSecret string `envconfig:"LINE_CHANNEL_SECRET" required:"true"`
}

var env Env

func Init() error {

	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env: %s", err)
		return err
	}

	return nil
}

func GetEnv() Env {
	return env
}