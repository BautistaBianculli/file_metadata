package application

import (
	"fmt"
	"github.com/BautistaBianculli/metadata_archivos/src/application/config"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/dependency"
	"github.com/BautistaBianculli/metadata_archivos/src/http/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApplication() {

	env := config.EnvVars()
	handlers := dependency.Start(env)

	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(cors.Default())

	router.Init(server, handlers)

	port := fmt.Sprintf(":%d", env.Port)

	err := server.Run(port)
	if err != nil {
		panic(err)
	}
}
