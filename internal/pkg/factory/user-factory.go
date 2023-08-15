package factory

import (
	"github.com/grootkng/clean-arch-golang/internal/domain/implementation"
	"github.com/grootkng/clean-arch-golang/internal/pkg/api/controller"
	"github.com/grootkng/clean-arch-golang/internal/pkg/repository/database"
)

func UserFactory() *controller.UserController {
	userRepository := database.NewUserRepository()
	userImplementation := implementation.NewUserImplementation(userRepository)
	return controller.NewUserController(userImplementation)
}
