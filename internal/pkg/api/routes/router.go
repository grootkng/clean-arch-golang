package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/grootkng/clean-arch-golang/docs"
	"github.com/grootkng/clean-arch-golang/internal/pkg/factory"
)

func SetupRouter() *gin.Engine {
	docs.SwaggerInfo.Title = "Clean Arch Golang"
	docs.SwaggerInfo.Description = "This is a Go server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()

	userController := factory.UserFactory()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("v1/users")
	{
		users.POST("", userController.Create)
		users.PUT("/:id", userController.UpdateBy)
		users.DELETE("/:id", userController.DeleteBy)
		users.GET("/:id", userController.FindBy)
		users.GET("", userController.FindAll)
	}

	return router
}
