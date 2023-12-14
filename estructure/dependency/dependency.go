package dependency

import (
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/BautistaBianculli/metadata_archivos/src/core/usecases"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/repository"
	"github.com/BautistaBianculli/metadata_archivos/src/http/entrypoints"
)

type HandleContainer struct {
	Files entrypoints.FileHandler
}

func Start(env config.Config) *HandleContainer {

	// restyClient := resty.New().R()

	//repository
	AWSsession := repository.NewAwsSession(env)
	fileClient := repository.NewFileRepository(AWSsession, env)

	//useCases
	useCases := usecases.NewFileCases(fileClient)

	return &HandleContainer{
		Files: &entrypoints.Files{FileUseCases: useCases},
	}

}
