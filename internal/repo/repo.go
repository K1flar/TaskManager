package repo

import "taskmanager/internal/domain"

type TaskRepo interface {
	GetAllTasks() ([]*domain.Task, error)
	GetTasksByAssignee(assigneeID uint32) ([]*domain.Task, error)
	GetTasksByAuthor(authorID uint32) ([]*domain.Task, error)
	AddTask(task *domain.Task) error
	Assign(id uint32, assignee string, assigneeID uint32) error
	Unassign(id uint32) error
	Resolve(id uint32) error

	GetUserByID(id uint32) (*domain.User, error)
	GetUserByTgID(tgID uint32) (*domain.User, error)
}
