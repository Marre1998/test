package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}
func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task

	result := r.db.First(&existingTask, id)
	if result.Error != nil {
		return task, result.Error
	}

	result = r.db.Model(&existingTask).Updates(task)
	if result.Error != nil {
		return task, result.Error
	}

	return existingTask, nil
}
func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	result := r.db.Where("id = ?", id).Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
