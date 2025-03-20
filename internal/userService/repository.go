package userService

import (
	"PetProject/internal/taskService"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
	GetAllUsers() ([]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository { return &userRepository{db: db} }

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	existingUser := User{}
	result := r.db.First(&existingUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	result = r.db.Model(&existingUser).Updates(user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return existingUser, nil
}

// filepath: c:\Users\Admin\Desktop\ПетПроект\test\internal\userService\repository.go
func (r *userRepository) DeleteUserByID(id uint) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Find the user
		var user User
		if err := tx.First(&user, id).Error; err != nil {
			return err
		}

        // Delete associated tasks
        if err := tx.Where("user_id = ?", id).Delete(&taskService.Task{}).Error; err != nil {
            return err
        }

        // Delete the user
        if err := tx.Delete(&user).Error; err != nil {
            return err
        }

        return nil
    })
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return []User{}, result.Error
	}
	return users, nil
}
