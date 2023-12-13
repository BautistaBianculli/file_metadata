package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AWSSecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" required:"true"`
	AWSSecretKeyID     string `envconfig:"AWS_ACCESS_KEY_ID" required:"true"`
	AWSRegion          string `envconfig:"AWS_REGION" required:"true"`
	AWSBucket          string `envconfig:"AWS_BUCKET" required:"true"`
	Port               int    `required:"true" default:"8080"`
}

func EnvVars() Config {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
