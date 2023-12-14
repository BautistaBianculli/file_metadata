package router

import (
	"github.com/BautistaBianculli/metadata_archivos/src/docs"
	"github.com/BautistaBianculli/metadata_archivos/src/estructure/dependency"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	apiUploadPrefix = "/api/file/"
)

func Init(server *gin.Engine, handlers *dependency.HandleContainer) {
	v1 := server.Group(apiUploadPrefix)

	initSwaggerEndpoint(server)
	initUploadFileEndpoint(v1, handlers)
	initGetAllFileEndpoint(v1, handlers)
	initGetFileEndpoint(v1, handlers)
	initDeleteFileEndpoint(v1, handlers)
}

func initSwaggerEndpoint(rg *gin.Engine) {
	docs.SwaggerInfo.BasePath = apiUploadPrefix
	rg.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func initUploadFileEndpoint(rg *gin.RouterGroup, handlers *dependency.HandleContainer) {
	rg.POST("upload", handlers.Files.UploadHandle)
}

func initGetAllFileEndpoint(rg *gin.RouterGroup, handlers *dependency.HandleContainer) {
	rg.GET("", handlers.Files.GetAllHandler)
}

func initGetFileEndpoint(rg *gin.RouterGroup, handlers *dependency.HandleContainer) {
	rg.GET("get/:filename", handlers.Files.GetOneHandler)
}

func initDeleteFileEndpoint(rg *gin.RouterGroup, handlers *dependency.HandleContainer) {
	rg.DELETE("delete/:filename", handlers.Files.DeleteOneHandler)
}
