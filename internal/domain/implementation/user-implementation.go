package implementation

import (
	"github.com/grootkng/clean-arch-golang/internal/domain/entity"
	"github.com/grootkng/clean-arch-golang/internal/domain/usecase"
)

type UserImplementation struct {
	Repository usecase.IUserUsecase
}

func NewUserImplementation(r usecase.IUserUsecase) *UserImplementation {
	return &UserImplementation{
		Repository: r,
	}
}

func (impl *UserImplementation) Create(user *entity.User) error {
	if err := impl.Repository.Create(user); err != nil {
		return err
	}

	return nil
}

func (impl *UserImplementation) FindAll(page int, pageSize int) ([]entity.User, error) {
	users, err := impl.Repository.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (impl *UserImplementation) FindBy(id int) (*entity.User, error) {
	user, err := impl.Repository.FindBy(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (impl *UserImplementation) UpdateBy(user *entity.User) error {
	if err := impl.Repository.UpdateBy(user); err != nil {
		return err
	}

	return nil
}

func (impl *UserImplementation) DeleteBy(id int) error {
	if err := impl.Repository.DeleteBy(id); err != nil {
		return err
	}

	return nil
}
