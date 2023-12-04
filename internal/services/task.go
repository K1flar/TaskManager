package services

import (
	"fmt"
	"taskmanager/internal/domain"
	"taskmanager/internal/repo"
)

type ServiceError struct {
	Err error
}

func (se *ServiceError) Error() string {
	return fmt.Sprintf("Service error: %s", se.Err.Error())
}

type TaskService struct {
	repo repo.TaskRepo
}

func NewTaskService(repo repo.TaskRepo) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (ts *TaskService) GetAll() ([]*domain.Task, error) {
	tasks, err := ts.repo.GetAllTasks()
	if err != nil {
		return nil, &ServiceError{Err: err}
	}
	return tasks, nil
}

func (ts *TaskService) CreateTask(about string, author string, authorID uint32) error {
	task := &domain.Task{About: about, Author: author, AuthorID: authorID}
	err := ts.repo.AddTask(task)
	if err != nil {
		return &ServiceError{Err: err}
	}
	return nil
}

func (ts *TaskService) Assign(id uint32)
