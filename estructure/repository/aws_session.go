package repository

import (
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func NewAwsSession(env config.Config) *session.Session {

	sesion, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AWSRegion),
	})

	if err != nil {
		log.Fatalf("Error al crear al sesion de AWS: %s", err)
	}
	return sesion
}
