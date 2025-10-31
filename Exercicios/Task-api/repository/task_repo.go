package repository

//acesso ao DB

import (
	"task-api/database"
	"task-api/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	FindAll() ([]models.Task, error)
	FindByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository() TaskRepository {
	return &taskRepo{db: database.DB}
}

func (r *taskRepo) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepo) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepo) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepo) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
