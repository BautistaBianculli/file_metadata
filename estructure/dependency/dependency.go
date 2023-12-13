package dependency

import (
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/BautistaBianculli/metadata_archivos/src/core/usecases"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/repository"
	"github.com/BautistaBianculli/metadata_archivos/src/http/entrypoints"
)

type HandleContainer struct {
	UploadFile entrypoints.Handler
}

func Start(env config.Config) *HandleContainer {

	// restyClient := resty.New().R()

	//repository
	AWSsession := repository.NewAwsSession(env)
	fileClient := repository.NewFileRepository(*AWSsession, env)

	//useCases
	useCases := usecases.NewUploadCases(fileClient)

	return &HandleContainer{
		UploadFile: &entrypoints.Upload{UploadUseCase: useCases},
	}

}
