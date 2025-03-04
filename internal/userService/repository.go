package userService

import "gorm.io/gorm"

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

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return []User{}, result.Error
	}
	return users, nil
}
