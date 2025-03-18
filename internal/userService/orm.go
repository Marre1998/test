package userService

import (
	"PetProject/internal/taskService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint               `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
