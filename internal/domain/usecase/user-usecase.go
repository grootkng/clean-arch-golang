package usecase

import (
	"github.com/grootkng/clean-arch-golang/internal/domain/entity"
)

type IUserUsecase interface {
	FindAll(page int, pageSize int) ([]entity.User, error)
	FindBy(id int) (*entity.User, error)
	UpdateBy(user *entity.User) error
	DeleteBy(id int) error
	Create(*entity.User) error
}
