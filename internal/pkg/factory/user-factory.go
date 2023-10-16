package factory

import (
	"github.com/grootkng/clean-arch-golang/internal/domain/implementation"
	"github.com/grootkng/clean-arch-golang/internal/pkg/api/controller"
	"github.com/grootkng/clean-arch-golang/internal/pkg/repository/database"
)

func UserFactory() *controller.UserController {
	db, err := database.Db()
	if err != nil {
		return nil
	}

	userRepository := database.NewUserRepository(db)
	userImplementation := implementation.NewUserImplementation(userRepository)
	return controller.NewUserController(userImplementation)
}
