package database

import (
	"github.com/grootkng/clean-arch-golang/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user *entity.User) error {
	ur.DB.Create(&user)
	return nil
}

func (ur *UserRepository) FindAll(page int, pageSize int) ([]entity.User, error) {
	var result []entity.User
	if err := ur.DB.Scopes(Paginate(page, pageSize)).Last(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (ur *UserRepository) FindBy(id int) (*entity.User, error) {
	var result entity.User
	ur.DB.Where("id = ?", id).First(&result)

	if result.Id == 0 {
		return nil, nil
	}

	return &result, nil
}

func (ur *UserRepository) UpdateBy(user *entity.User) error {
	ur.DB.Model(&user).UpdateColumns(entity.User{Name: user.Name, Age: user.Age, Gender: user.Gender})
	return nil
}

func (ur *UserRepository) DeleteBy(id int) error {
	user := entity.User{}
	user.Id = id
	ur.DB.Delete(&user)
	return nil
}
