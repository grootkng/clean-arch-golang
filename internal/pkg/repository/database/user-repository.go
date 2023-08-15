package database

import (
	"github.com/grootkng/clean-arch-golang/internal/domain/entity"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (impl *UserRepository) Create(user *entity.User) error {
	db, err := Db()
	if err != nil {
		return err
	}

	db.Create(&user)
	return nil
}

func (impl *UserRepository) FindAll() ([]entity.User, error) {
	db, err := Db()
	if err != nil {
		return []entity.User{}, err
	}

	var result []entity.User
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (impl *UserRepository) FindBy(id int) (*entity.User, error) {
	db, err := Db()
	if err != nil {
		return nil, err
	}

	var result entity.User
	db.Where("id = ?", id).First(&result)

	if result.Id == 0 {
		return nil, nil
	}

	return &result, nil
}

func (impl *UserRepository) UpdateBy(user *entity.User) error {
	db, err := Db()
	if err != nil {
		return err
	}

	db.Model(&user).UpdateColumns(entity.User{Name: user.Name, Age: user.Age, Gender: user.Gender})
	return nil
}

func (impl *UserRepository) DeleteBy(id int) error {
	db, err := Db()
	if err != nil {
		return err
	}

	user := entity.User{}
	user.Id = id
	db.Delete(&user)
	return nil
}
