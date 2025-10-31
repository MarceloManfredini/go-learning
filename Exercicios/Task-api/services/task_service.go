package services

//lógica de negócio

import (
	"errors"
	"task-api/models"
	"task-api/repository"
)

type TaskService interface {
	CreateTask(t *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(id uint, t *models.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(t *models.Task) error {
	return s.repo.Create(t)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.FindAll()
}

func (s *taskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.FindByID(id)
}

func (s *taskService) UpdateTask(id uint, t *models.Task) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	existing.Title = t.Title
	existing.Detail = t.Detail
	existing.Done = t.Done
	return s.repo.Update(existing)
}

func (s *taskService) DeleteTask(id uint) error {
	// optional: verify existence
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("task not found")
	}
	return s.repo.Delete(id)
}
