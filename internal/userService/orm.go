// filepath: c:\Users\Admin\Desktop\ПетПроект\test\internal\userService\orm.go
package userService

import (
    "PetProject/internal/taskService"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email    string             `json:"email"`
    Password string             `json:"password"`
    Tasks    []taskService.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}