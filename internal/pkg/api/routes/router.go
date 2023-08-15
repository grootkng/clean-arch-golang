package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grootkng/clean-arch-golang/internal/pkg/factory"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	userController := factory.UserFactory()

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
